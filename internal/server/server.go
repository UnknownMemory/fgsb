package server

import (
	"embed"
	"fgsb/internal/server/handler"
	"io/fs"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Server struct {
	Theme string `json:"theme"`
	Port int `json:"port"`
}

var Assets fs.FS
var Templates embed.FS

func (s *Server) Run() {
	handler.Templates = Templates
	handler.Theme = s.Theme
	
	mux := http.NewServeMux()

	assetsFS := http.FileServer(http.FS(Assets))
	mux.Handle("/assets/", http.StripPrefix("/assets/", disableDirList(assetsFS)))

	themeFS := http.FileServer(http.Dir("./themes/"+s.Theme))
	mux.Handle("/themes/"+s.Theme+"/", http.StripPrefix("/themes/"+s.Theme+"/", disableDirList(themeFS)))

	mux.HandleFunc("/{$}", handler.Root)
	mux.HandleFunc("/admin/edit-scoreboard", handler.EditScoreboard)
	
	mux.HandleFunc("/api/v1/scoreboard/events", handler.SSEEvents)
	mux.HandleFunc("POST /api/v1/scoreboard/update", handler.SSEUpdate)

	addr := ":" + strconv.Itoa(s.Port)
	http.ListenAndServe(addr, mux)
}


func (s *Server) Open(url string) error {
	addr := ":" + strconv.Itoa(s.Port)

	cmd := exec.Command("cmd", "/c", "start", "http://localhost"+addr+url)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}


func disableDirList(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") {
            http.NotFound(w, r)
            return
        }

        handler.ServeHTTP(w, r)
    })
}
