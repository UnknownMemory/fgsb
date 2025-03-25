package handler

import (
	"embed"
	"html/template"
	"net/http"
)

var Templates embed.FS

type TemplateData struct {
	Title string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	t, err := template.ParseFS(Templates, "web/templates/include/base.html" , "web/templates/pages/" + tmpl + ".html")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	err = t.Execute(w, data)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
