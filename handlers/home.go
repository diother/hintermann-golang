package handlers

import (
	"text/template"
	"log"
	"net/http"
)

var counter int = 1

func arr(els ...interface{}) []interface{} {
	return els
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base").Funcs(template.FuncMap {
		"arr": arr,
	})
	tmpl = template.Must(tmpl.ParseGlob("views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/components/*.html"))

	err := tmpl.ExecuteTemplate(w, "home", nil)
	if err != nil {
        log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base").Funcs(template.FuncMap {
		"arr": arr,
	})
	tmpl = template.Must(tmpl.ParseGlob("views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/components/*.html"))

	err := tmpl.ExecuteTemplate(w, "contact", nil)
	if err != nil {
        log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
