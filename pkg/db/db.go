package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"kanban/pkg/config"
)

func NewDBConnection(cfg *config.Config) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	return dbpool, nil
}
