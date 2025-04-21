package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)


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
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		http.Error(w, "Unable to parse multipart", http.StatusBadRequest)
		return
	}

	data := make(map[string]string, len(r.MultipartForm.Value))
	for key, values := range r.MultipartForm.Value {
		data[key] = values[0]
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
