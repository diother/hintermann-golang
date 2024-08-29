package handlers

import (
	"text/template"
	"log"
	"net/http"
)

type PageData struct {
    Title string
    Content string
    Count int
}

var counter int = 1

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "Welcome",
        Content: "This is the home page.",
        Count: counter,
    }

    tmpl := template.Must(template.ParseFiles(
        "views/home.tmpl", 
        "views/head.tmpl", 
        "views/foot.tmpl", 
        "views/components/header.tmpl", 
        "views/components/icons.tmpl",
    ))

    err := tmpl.ExecuteTemplate(w, "home", data)
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/home.tmpl"))

    data := PageData{
        Title:   "Welcome",
        Content: "This is the about page.",
        Count: counter,
    }

    err := tmpl.ExecuteTemplate(w, "home", data)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func CountHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.NotFound(w, r)
        return
    }
    tmpl := template.Must(template.ParseFiles("views/components/header.tmpl"))

    err := tmpl.ExecuteTemplate(w, "header", nil)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
