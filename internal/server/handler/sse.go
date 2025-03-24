package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FormData struct {
	Player1 string
	Score1 string
	Player2 string
	Score2 string
}

var chMessage chan []byte

func SSEEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	
	flusher, ok := w.(http.Flusher)
	if !ok {
        http.Error(w, "Streaming not supported.", http.StatusNotImplemented)
        return
	}

	chMessage = make(chan []byte)
	
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

func SSEUpdate(w http.ResponseWriter, r *http.Request) {
	data := FormData{
		Player1: r.FormValue("name1"),
		Score1: r.FormValue("score1"),
		Player2: r.FormValue("name2"),
		Score2: r.FormValue("score2"),
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	chMessage <- jsonBytes
}
