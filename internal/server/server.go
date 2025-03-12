package server

import (
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
	
	http.ListenAndServe(s.addr, mux)
}


