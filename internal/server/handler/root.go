package handler

import (
	"html/template"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	base_theme, err := template.ParseFS(Templates, "web/templates/include/base_theme.html")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	t, err := template.Must(base_theme.Clone()).ParseFiles("./themes/"+Theme+"/index.html")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	err = t.Execute(w, nil)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
