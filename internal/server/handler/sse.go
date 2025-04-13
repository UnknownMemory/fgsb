package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

type FormData struct {
	Player1 string
	Score1 string
	Player2 string
	Score2 string
}

var (
	idCounter uint64
	clients map[string]chan []byte = make(map[string]chan []byte)
	mu sync.Mutex
)

func generateID() string {
	return strconv.FormatUint(atomic.AddUint64(&idCounter, 1), 10)
}

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

	clientID := generateID()
	chMessage := make(chan []byte)

	mu.Lock()
	clients[clientID] = chMessage
	mu.Unlock()
	
	for {
		select {
		case msg := <-chMessage:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()

		case <-r.Context().Done():
			mu.Lock()
			delete(clients, clientID)
			mu.Unlock()
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

	mu.Lock()
	defer mu.Unlock()

	for _, chMessage := range clients {
		chMessage <- jsonBytes
	}

}
