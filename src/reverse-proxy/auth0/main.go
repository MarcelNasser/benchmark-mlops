package main

import (
    "log"
    "net/http"
 	"github.com/joho/godotenv"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/platform/authenticator"
	"github.com/MarcelNasser/benchmark-mlops/reverse-proxy/auth0/platform/router"
)


func main() {
    //check .env present
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
    //auth
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	//start Main App server
    log.Fatal(http.ListenAndServe(":8080", router.App(auth)))
}