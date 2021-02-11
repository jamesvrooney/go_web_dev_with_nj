package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>This is a test</h1>")
	})

	http.ListenAndServe(":8181", nil)
}
