package handler

import (
	"fmt"
	"net/http"
)

var chMessage chan string

func ScoreboardEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	
	flusher, ok := w.(http.Flusher)
	if !ok {
        http.Error(w, "Streaming not supported.", http.StatusNotImplemented)
        return
	}

	chMessage = make(chan string)
	
	for {
		select {
		case msg := <-chMessage:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()

		case <-r.Context().Done():
			close(chMessage)
			return
		}
		
	}
}

func ScoreboardUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	chMessage <-"{'ab': 2}"

}
