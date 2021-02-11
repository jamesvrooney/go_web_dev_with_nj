package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello This handles requests to the /hello endpoint
type Hello struct {
	l *log.Logger
}

// NewHello Returns a Hello struct
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Hello %s", d)
}
