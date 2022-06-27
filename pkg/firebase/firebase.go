package firebase

// FirestoreUser is a user in Firestore
type FirestoreUser struct {
	Email string `firestore:"email" json:"email"`
	ID string `firestore:"id" json:"id"`
	DisplayName string `firestore:"displayName" json:"displayName"`
	Username string `firestore:"username" json:"username"`
}

// APIToken is the Firebase API token.
type APIToken struct {
	CreatedAt string `firestore:"createdAt" json:"createdAt"`
	ExpiredAt string `firestore:"expiredAt" json:"expiredAt"`
	Token string `firestore:"token" json:"token"`
	ExpiresIn int `firestore:"expiresIn" json:"expiresIn"`
}