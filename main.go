package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/JohnVicke/klask-backend/cmd/healthcheck"
	"github.com/JohnVicke/klask-backend/cmd/helloworld"
)

func main() {
	port := "8080"

	funcframework.RegisterHTTPFunction("/hello-world", helloworld.HelloWorld)
	funcframework.RegisterHTTPFunction("/healthcheck", healthcheck.HealthCheck)

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}


	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}