package helloworld

import (
	"net/http"
)

// HelloWorld is a function that returns a string.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}