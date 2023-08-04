package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
    "fmt"
	"github.com/gin-gonic/gin"
)


func Handler(remote *url.URL) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = ctx.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            req.URL.Path = ctx.Param("proxyPath")
        }
        proxy.ServeHTTP(ctx.Writer, ctx.Request)
    }
}

func Static(remote *url.URL) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = ctx.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            dir := ctx.Params.ByName("dir")
            file := ctx.Params.ByName("file")
            if &dir != nil && &file != nil {
              req.URL.Path = fmt.Sprintf("/static-files/static/%s/%s", ctx.Params.ByName("dir"), ctx.Params.ByName("file"))
            } else {
              req.URL.Path = fmt.Sprintf("/static-files/%s", ctx.Params.ByName("file"))
            }
        }
        proxy.ServeHTTP(ctx.Writer, ctx.Request)
    }
}

func RunApi(remote *url.URL) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = ctx.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            req.URL.Path = fmt.Sprintf("/ajax-api/2.0/mlflow/%s/%s", ctx.Params.ByName("dir"), ctx.Params.ByName("file"))
            req.Method = ctx.Request.Method
        }
        proxy.ServeHTTP(ctx.Writer, ctx.Request)
    }
}

func ArtifactApi(remote *url.URL) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = ctx.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            req.URL.Path = fmt.Sprintf("/get-artifact")
            req.Method = ctx.Request.Method
            paramPairs := ctx.Request.URL.Query()
            for key, value := range paramPairs {
                req.URL.Query().Add(key, value[0])
            }
        }
        proxy.ServeHTTP(ctx.Writer, ctx.Request)
    }
}
