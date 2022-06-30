package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/JohnVicke/klask-backend/cmd/healthcheck"
	"github.com/joho/godotenv"
)

func init() {	

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error [godotenv] loading .env file: %v\n", err)
	}
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