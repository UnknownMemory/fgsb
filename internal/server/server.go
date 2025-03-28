package server

import (
	"embed"
	"fgsb/internal/server/handler"
	"io/fs"
	"net/http"
	"os/exec"
	"strconv"
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
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(Assets))))
	mux.Handle("/themes/"+s.Theme+"/", http.StripPrefix("/themes/"+s.Theme+"/", http.FileServer(http.Dir("./themes/"+s.Theme))))

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
