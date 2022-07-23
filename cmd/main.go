package main

import (
	"SocialConnections/config"
	"SocialConnections/pkg/postgres"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfdFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfdFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	
	psqlDB, err := postgres.InitPsqlDB(context.Background(), cfg)
	if err != nil {
		log.Fatalf("PostgreSQL init error: %s", err)
	} else {
		log.Println("PostgreSQL connected, status: %#v", psqlDB.Stats())
	}

	defer func(psqlDB *sqlx.DB) {
		err = psqlDB.Close()
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("PostgreSQL closed properly")
		}
	}(psqlDB)

	
}