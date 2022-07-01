package firebase

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

// FirestoreUser is a user in Firestore
type FirestoreUser struct {
	Email string `firestore:"email" json:"email"`
	ID string `firestore:"id" json:"id"`
	DisplayName string `firestore:"displayName" json:"displayName"`
	Username string `firestore:"username" json:"username"`
	UpdatedAt string `firestore:"updatedAt" json:"updatedAt"`
}

// APIToken is the Firebase API token.
type APIToken struct {
	CreatedAt string `firestore:"createdAt" json:"createdAt"`
	ExpiredAt string `firestore:"expiredAt" json:"expiredAt"`
	Token string `firestore:"token" json:"token"`
	ExpiresIn int `firestore:"expiresIn" json:"expiresIn"`
}

func init() {
	godotenv.Load()
	log.Println("Initialized firebase with .env file")
}

// InitializeFirebaseWithEnv initializes Firebase with the environment variables.
func InitializeFirebaseWithEnv () (*firestore.Client, error) {
	ctx := context.Background()
	projectID := os.Getenv("FIREBASE_PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
		return nil, err
	}

	return client, nil
}