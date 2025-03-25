package server

import (
	"embed"
	"fgsb/internal/server/handler"
	"io/fs"
	"net/http"
	"strconv"
)

type Server struct {
	addr string
}

func NewServer(addr int) *Server {
	cAddr := strconv.Itoa(addr)
	return &Server{addr: ":"+cAddr}
}

var Assets fs.FS
var Templates embed.FS

func (s *Server) Run() {
	handler.Templates = Templates
	
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(Assets))))

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/admin/edit-scoreboard", handler.EditScoreboard)
	
	mux.HandleFunc("/api/scoreboard/events", handler.SSEEvents)
	mux.HandleFunc("POST /api/scoreboard/update", handler.SSEUpdate)

	
	http.ListenAndServe(s.addr, mux)
}


