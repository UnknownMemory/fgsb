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
	mux.Handle("/themes/"+s.Theme, http.StripPrefix("/themes/"+s.Theme, http.FileServer(http.Dir("./themes/"+s.Theme))))

	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/admin/edit-scoreboard", handler.EditScoreboard)
	
	mux.HandleFunc("/api/v1/scoreboard/events", handler.SSEEvents)
	mux.HandleFunc("POST /api/v1/scoreboard/update", handler.SSEUpdate)

	addr := ":" + strconv.Itoa(s.Port)
	http.ListenAndServe(addr, mux)
}


func (s *Server) Open(url string) error {
    var cmd string
    var args []string
	
	addr := ":" + strconv.Itoa(s.Port)

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start", "http://localhost"+addr+url}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, "http://localhost"+addr+url)
    return exec.Command(cmd, args...).Start()
}
