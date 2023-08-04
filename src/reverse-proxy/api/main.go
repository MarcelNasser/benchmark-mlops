package main

import (
	"os"
    "log"
    "net/http"
	"net/url"
	"fmt"
    "strings"
	"net/http/httputil"
 	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

)


func main() {
    gin.SetMode(gin.ReleaseMode)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
    log.Fatal(http.ListenAndServe(":9000", Router()))
}


func Router() *gin.Engine {
	router := gin.Default()
    remote, err := url.Parse(os.Getenv("MLFLOW_HOST"))
	if err!=nil {
		panic(err)
	}
    //mlflow / external Api
	r1 := [] string{"mlflow"}
	r2 := [] string{"mlflow-artifacts", "artifacts"}
	r3 := [] string{"mlflow-artifacts", "artifacts","model"}
    router.GET("/api/2.0/mlflow/:dir/:file", RestApi(remote, r1))
    router.POST("/api/2.0/mlflow/:dir/:file", RestApi(remote, r1))
    router.GET("/api/2.0/mlflow-artifacts/artifacts/:xp/:artifact/artifacts/:file", RestApi(remote, r2))
    router.PUT("/api/2.0/mlflow-artifacts/artifacts/:xp/:artifact/artifacts/:file", RestApi(remote, r2))
    router.GET("/api/2.0/mlflow-artifacts/artifacts/:xp/:artifact/artifacts/model/:file", RestApi(remote, r3))
    router.PUT("/api/2.0/mlflow-artifacts/artifacts/:xp/:artifact/artifacts/model/:file", RestApi(remote, r3))
	return router
}

func RestApi(remote *url.URL, routes []string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        proxy := httputil.NewSingleHostReverseProxy(remote)
        proxy.Director = func(req *http.Request) {
            req.Header = ctx.Request.Header
            req.Host = remote.Host
            req.URL.Scheme = remote.Scheme
            req.URL.Host = remote.Host
            req.Method = ctx.Request.Method
            if routes[0]=="mlflow"{
                dir := ctx.Params.ByName("dir")
                file := ctx.Params.ByName("file")
                req.URL.Path = fmt.Sprintf("/api/2.0/mlflow/%s/%s", dir, file)
            }
            if routes[0]=="mlflow-artifacts"{
                xp := ctx.Params.ByName("xp")
                artifact := ctx.Params.ByName("artifact")
                file := ctx.Params.ByName("file")
                req.URL.Path = fmt.Sprintf("/api/2.0/mlflow-artifacts/artifacts/%s/%s/%s/%s", xp, artifact, strings.Join(routes[1:],"/"), file)
            }
        }
        proxy.ServeHTTP(ctx.Writer, ctx.Request)
    }
}