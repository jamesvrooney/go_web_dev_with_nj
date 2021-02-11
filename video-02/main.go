package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jamesvrooney/example/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(logger)

	http.ListenAndServe(":8181", nil)
}
