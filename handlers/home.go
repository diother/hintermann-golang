package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
    Title string
    Content string
    Counter int
}

var counter int = 1

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // ignore requests for styles and favicon
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    tmpl := template.Must(template.ParseFiles("views/index.html"))

    counter++ 
    data := PageData{
        Title:   "Welcome",
        Content: "This is the home page.",
        Counter: counter,
    }

    err := tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
