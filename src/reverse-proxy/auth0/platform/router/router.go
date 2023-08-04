package router

import (
	"encoding/gob"
	"os"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/url"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/platform/authenticator"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/platform/middleware"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/web/app/callback"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/web/app/home"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/web/app/login"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/web/app/logout"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/web/app/proxy"
)

// New registers the routes and returns the router.
func App(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()
    remote, err := url.Parse(os.Getenv("MLFLOW_HOST"))
	if err!=nil {
		panic(err)
	}
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))
    //auth0 / static
	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")
	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
    router.GET("/logout", logout.Handler)
	router.GET("/callback", callback.Handler(auth))
    //mlflow / redirection to root
	router.GET("/root", middleware.IsAuthenticated, proxy.Handler(remote))
    //mlflow / static
    router.GET("/static-files/:file", middleware.IsAuthenticated, proxy.Static(remote))
    router.GET("/static-files/static/:dir/:file", middleware.IsAuthenticated, proxy.Static(remote))
    //mlflow / internal APi
    router.GET("/ajax-api/2.0/mlflow/:dir/:file", middleware.IsAuthenticated, proxy.RunApi(remote))
    router.POST("/ajax-api/2.0/mlflow/:dir/:file", middleware.IsAuthenticated, proxy.RunApi(remote))
    router.GET("/get-artifact", middleware.IsAuthenticated, proxy.ArtifactApi(remote))
    router.POST("/get-artifact", middleware.IsAuthenticated, proxy.ArtifactApi(remote))
	return router
}
