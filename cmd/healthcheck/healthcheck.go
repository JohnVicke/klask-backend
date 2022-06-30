package healthcheck

import (
	"log"
	"net/http"
)

func init() {
	log.Println("Healthcheck initialized")
}

// HealthCheck is a function that returns a string.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!!"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}