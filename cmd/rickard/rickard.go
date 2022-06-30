package rickard


import (
	"log"
	"net/http"
)

func init() {
	log.Println("Healthcheck initialized")
}

// Rickard is a function that returns a string.
func Rickard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Rickard"))
	w.WriteHeader(http.StatusOK)
}
