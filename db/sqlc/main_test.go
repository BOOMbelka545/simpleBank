package db

import (
	"context"
	"os"
	"simpleBank/config"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
