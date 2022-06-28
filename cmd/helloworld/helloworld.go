package helloworld

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloWorld is a function that returns a string.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, World!")
	w.Write([]byte("Hello, World!"))
}