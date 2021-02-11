package main

import (
	"net/http"

	"github.com/jamesvrooney/go_web_dev_with_nj/handlers"
)

func main() {

	hh := handlers.NewHelloHandler()

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)

	http.ListenAndServe(":9876", hh)
}
