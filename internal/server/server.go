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

	mux.HandleFunc("/", handler.Root)
	
	http.ListenAndServe(s.addr, mux)
}


