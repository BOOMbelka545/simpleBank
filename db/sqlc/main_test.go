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
var testDB *pgx.Conn

func TestMain(m *testing.M) {
	var err error

	testDB, err = pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
