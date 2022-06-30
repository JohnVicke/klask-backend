package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/JohnVicke/klask-backend/cmd/healthcheck"
)

func init() {	
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Loaded .env file")
}

func main() {
	funcframework.RegisterHTTPFunction("/healthcheck", healthcheck.HealthCheck)

	port := "8080"

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}


	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}