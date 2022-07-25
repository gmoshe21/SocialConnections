package postgres

import (
	"SocialConnections/config"
	"SocialConnections/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

const (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS social_connections (
		id SERIAL PRIMARY KEY,
		sender VARCHAR(100),
		recipient VARCHAR(100),
		number_of_communications INTEGER
	  );`
	queryInsert = `INSERT INTO social_connections ( sender, recipient, number_of_communications)
					VALUES ($1, $2, 0);`
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

	database.MustExec(queryCreateTable)

	data, err := ioutil.ReadFile("./graf.json")
    if err != nil {
      fmt.Print(err)
    }

	var graf []models.Connections

	err = json.Unmarshal(data, &graf)
	if err != nil {
		return nil, err
	}

	for _, val := range graf {
		_, err = database.ExecContext(context.Background(), queryInsert, val.Sender, val.Sender)
		if err != nil {
			return nil, err
		}
	}

	return database, nil
}