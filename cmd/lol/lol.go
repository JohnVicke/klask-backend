package lol

import (
	"log"
	"net/http"
)

func init() {
	log.Println("Healthcheck initialized")
}

// Lol is a function that returns a string.
func Lol(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Lol"))
	w.WriteHeader(http.StatusOK)
}
