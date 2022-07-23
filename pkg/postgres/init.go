package postgres

import (
	"SocialConnections/config"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InitPsqlDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	database, err := sqlx.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}