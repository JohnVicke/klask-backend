package healthcheck

import (
	"log"
	"net/http"
	"os"
)

func init() {
	log.Println("Healthcheck initialized")
}

// HealthCheck is a function that returns a string.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var1 := os.Getenv("VAR1")
	w.Write([]byte("Please ork!!!" + var1))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}