package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/handlers"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
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

	// Logger
	logger, loggerInitErr := zap.NewDevelopment()
	if loggerInitErr != nil {
		os.Exit(1)
	}

	// Config file
	apiCfg := handlers.ApiConfig{
		Db:          database.New(db),
		Platform:    os.Getenv("PLATFORM"),
		TokenSecret: os.Getenv("TOKEN_SECRET"),
		Logger:      logger,
	}

	// New Mux
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", apiCfg.Register)
	mux.HandleFunc("POST /login", apiCfg.Login)

	// New Http Server
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
