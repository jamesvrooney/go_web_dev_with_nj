package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(logger)

	http.ListenAndServe(":8181", nil)
}
