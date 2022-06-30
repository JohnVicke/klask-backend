package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/JohnVicke/klask-backend/cmd/healthcheck"
	"github.com/JohnVicke/klask-backend/cmd/rickard"
)

func init() {	
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Loaded .env file")
}

func main() {
	funcframework.RegisterHTTPFunction("/healthcheck", healthcheck.HealthCheck)
	funcframework.RegisterHTTPFunction("/rickard", rickard.Rickard)

	port := "8080"

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}


	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}