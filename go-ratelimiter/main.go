package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	rl := NewRateLimiter(1, 5)
	http.HandleFunc("/", rl.Limit(handler))
	err := http.ListenAndServe(
		":8080",
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to listen on port 8080, %s", err)
	}
}

type Message struct {
	Status string
	Body   string
}

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Status: "OK",
		Body:   "Hi! from server...",
	}
	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		return
	}
}
