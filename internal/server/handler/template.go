package handler

import (
	"html/template"
	"net/http"
)

type TemplateData struct {
	Title string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	t, err := template.ParseFiles("web/templates/" + tmpl + ".html")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	err = t.Execute(w, data)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
