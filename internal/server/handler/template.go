package handler

import (
	"embed"
	"html/template"
	"net/http"
)

var Templates embed.FS
var Theme string

func renderTemplate[T any](w http.ResponseWriter, tmpl string, data T) {
	t, err := template.ParseFS(Templates, "web/templates/include/base.html", "web/templates/pages/"+tmpl+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
