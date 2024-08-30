package main

import (
	"log"
	"net/http"

	"github/diother/hintermann-go/handlers"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/contact", handlers.ContactHandler)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

