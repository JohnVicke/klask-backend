package createuser

import (
	"log"
	"net/http"

	firestore "cloud.google.com/go/firestore"
)

func init () {
	log.Println("Initializing createuser")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating user")
	ctx := r.Context()
	client, err := firestore.NewClient(ctx, "functions-framework-go")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc("user1").Set(ctx, map[string]interface{}{
		"name": "John Doe",
		"age":   30,
	})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
}