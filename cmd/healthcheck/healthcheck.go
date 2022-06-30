package healthcheck

import (
	"net/http"
	"os"
)

// HealthCheck is a function that returns a string.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var1 := os.Getenv("VAR1")
	w.Write([]byte("Please work!!" + var1))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}