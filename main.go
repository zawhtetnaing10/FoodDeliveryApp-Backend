package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/handlers"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
)

func main() {

	// Load env file
	if err := godotenv.Load(); err != nil {
		os.Exit(1)
	}

	dbUrl := os.Getenv("DATABASE_URL")

	// Connect to DB
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Printf("Unable to connect to db %v", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Ping the connection
	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		fmt.Printf("DB ping error %v\n", pingErr)
	}

	// Create db queries instance
	Db := database.New(pool)

	fmt.Printf("Successfully connected to database.")

	// Logger
	logger, loggerInitErr := zap.NewDevelopment()
	if loggerInitErr != nil {
		os.Exit(1)
	}

	// Config file
	apiCfg := handlers.ApiConfig{
		Pool:        pool,
		Db:          Db,
		Platform:    os.Getenv("PLATFORM"),
		TokenSecret: os.Getenv("TOKEN_SECRET"),
		Logger:      logger,
	}

	// New Mux
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/register", apiCfg.Register)
	mux.HandleFunc("POST /api/login", apiCfg.Login)
	mux.HandleFunc("POST /api/forget-password-check-email", apiCfg.CheckEmail)
	mux.HandleFunc("POST /api/forget-password", apiCfg.ForgetPassword)

	mux.HandleFunc("GET /api/get-all-restaurants", apiCfg.GetAllRestaurantWithCategories)
	mux.HandleFunc("GET /api/get-restaurant-details/{restaurant_id}", apiCfg.GetRestaurantDetails)

	mux.HandleFunc("POST /api/add-delivery-address-and-payment-method", apiCfg.AddDeliveryAddressAndPaymentMethod)
	mux.HandleFunc("GET /api/get-delivery-addresses-and-payment-methods", apiCfg.GetDeliveryAddressAndPaymentMethodForUser)

	mux.HandleFunc("POST /api/submit-order", apiCfg.SubmitOrder)
	mux.HandleFunc("GET /api/get-orders-for-user", apiCfg.GetOrdersForUser)

	// New Http Server
	server := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		apiCfg.LogError("Cannot serve endpoints", err)
		os.Exit(1)
	}
}
