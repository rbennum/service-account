package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rbennum/service-account/utils/config"
)

func New() (*pgxpool.Pool, error) {
	DbString := config.GetConfig().DBConnection
	db, err := pgxpool.New(context.Background(), DbString)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return db, nil
}
