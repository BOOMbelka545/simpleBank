package main

import (
	"context"
	"simpleBank/api"
	db "simpleBank/db/sqlc"
	"simpleBank/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
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
