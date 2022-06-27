package helloworld

import (
	"fmt"
	"net/http"
)

// HelloWorld is a function that returns a string.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	w.Write([]byte("Hello, World!"))
}