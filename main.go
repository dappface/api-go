package main

import (
	"log"
	"net/http"

	"github.com/dappface/api-go/registry"
)

const port = "8080"

func main() {
	registry := registry.NewRegistry()

	h := registry.NewHandler()
	http.Handle("/", h)

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
