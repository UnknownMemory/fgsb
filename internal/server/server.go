package server

import (
	"fgsb/internal/server/handler"
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

func (s *Server) Run() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/admin/edit-scoreboard", handler.EditScoreboard)
	
	mux.HandleFunc("/api/scoreboard/events", handler.SSEEvents)
	mux.HandleFunc("POST /api/scoreboard/update", handler.SSEUpdate)

	
	http.ListenAndServe(s.addr, mux)
}


