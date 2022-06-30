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
	w.Write([]byte("Hello, World from HealthCheck!"))
	w.WriteHeader(http.StatusOK)
}