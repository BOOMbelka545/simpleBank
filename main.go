package main

import (
	"context"
	"simpleBank/api"
	"simpleBank/config"
	db "simpleBank/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

func main() {
	conn, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start service: ", err)
	}

}
