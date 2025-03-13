package handler

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	data := &TemplateData{
		Title: "Scoreboard / FGSB",
	}
	renderTemplate(w, "index", data)
}
