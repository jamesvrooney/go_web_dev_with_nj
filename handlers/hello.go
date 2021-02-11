package handlers

import (
	"fmt"
	"net/http"
)

// Hello handler for requests to /hello
type Hello struct {
}

// NewHelloHandler jkjk
func NewHelloHandler() *Hello {
	return &Hello{}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello James</h1>")
}
