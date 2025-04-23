package main

import (
	"log"
	"net/http"

	"github.com/diother/hintermann-golang/internal/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	handler := handlers.NewHandler()

	http.HandleFunc("/", handler.HandleHome)
	// http.HandleFunc("/contact", handlers.ContactHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
