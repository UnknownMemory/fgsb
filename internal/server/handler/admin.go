package handler

import "net/http"

type AdminTemplateData struct {
	Title     string
	Countries map[string]string
}

func EditScoreboard(w http.ResponseWriter, r *http.Request) {

	data := AdminTemplateData{
		Title:     "Edit Scoreboard / FGSB",
		Countries: Countries,
	}

	renderTemplate(w, "admin/edit_scoreboard", data)
}
