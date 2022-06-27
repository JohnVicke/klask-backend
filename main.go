package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/JohnVicke/klask-backend/cmd/helloworld"
)

func main() {
	port := "8080"

	funcframework.RegisterHTTPFunction("/hello-world", helloworld.HelloWorld)

	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}


	if err:= funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

	fmt.Printf("Listening on port %s\n", port)
}