package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/diother/hintermann-golang/internal/helpers"
)

type Handler struct {
	tmpl *template.Template
}

func NewHandler() *Handler {
	tmpl := template.New("base").Funcs(template.FuncMap{
		"slice": helpers.SliceHelper,
		"attr":  helpers.AttrHelper,
	})
	tmpl = template.Must(tmpl.ParseGlob("internal/views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("internal/views/components/*.html"))

	return &Handler{tmpl}
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "home", nil); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
