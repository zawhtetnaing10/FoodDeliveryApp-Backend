package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// Load env file
	godotenv.Load()

	dbUrl := os.Getenv("DB_URL")

	// Open DB Connection
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Printf("DB connection error %v", err)
		os.Exit(1)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Printf("DB ping error %v", pingErr)
		os.Exit(1)
	}

	fmt.Printf("Successfully connected to database.")

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
