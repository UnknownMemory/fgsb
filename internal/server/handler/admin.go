package handler

import "net/http"

func EditScoreboard(w http.ResponseWriter, r *http.Request) {
	data := &TemplateData{
		Title: "Edit Scoreboard / FGSB",
	}

	renderTemplate(w, "admin/edit_scoreboard", data)

}
