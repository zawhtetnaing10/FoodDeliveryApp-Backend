package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
	"go.uber.org/zap"
)

type ApiConfig struct {
	Pool        *pgxpool.Pool
	Platform    string
	TokenSecret string
	Db          *database.Queries
	Logger      *zap.Logger
}

func (cfg *ApiConfig) LogError(message string, err error) {
	cfg.Logger.Error(message, zap.Error(err))
}
