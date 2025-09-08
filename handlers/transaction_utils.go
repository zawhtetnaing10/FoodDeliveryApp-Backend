package handlers

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/jackc/pgx/v5"
	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/internal/database"
)

func (cfg *ApiConfig) WithTransaction(ctx context.Context, fn func(q *database.Queries) error) (err error) {
	tx, err := cfg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}

	qtx := cfg.Db.WithTx(tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			err = fmt.Errorf("db operation failed: %v\n%s", p, debug.Stack())
			return
		} else if err != nil {
			_ = tx.Rollback(ctx)
			return
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				cfg.LogError("error committing transaction", err)
			}
		}
	}()

	err = fn(qtx)

	return err
}
