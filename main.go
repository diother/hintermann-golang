package main

import (
	"log"
	"net/http"

	"github/diother/hintermann-go/handlers"
)

func main() {
    // Define routes
    http.HandleFunc("/", handlers.HomeHandler)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

