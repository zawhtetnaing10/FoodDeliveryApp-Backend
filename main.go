package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// New Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthCheckHandler)

	// New Http Server
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

type StatusResponse struct {
	Status string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := StatusResponse{Status: "ok"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
