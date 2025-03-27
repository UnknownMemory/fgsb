package server

import (
	"embed"
	"fgsb/internal/server/handler"
	"io/fs"
	"net/http"
	"os/exec"
	"runtime"
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
	
	mux.HandleFunc("/api/v1/scoreboard/events", handler.SSEEvents)
	mux.HandleFunc("POST /api/v1/scoreboard/update", handler.SSEUpdate)

	
	http.ListenAndServe(s.addr, mux)
}


func (s *Server) Open(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start", "http://localhost"+s.addr+url}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, "http://localhost"+s.addr+"/"+url)
    return exec.Command(cmd, args...).Start()
}
