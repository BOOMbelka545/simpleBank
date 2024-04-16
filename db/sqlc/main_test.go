package db

import (
	"context"
	"os"
	"simpleBank/config"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

var testQueriesDB *Queries
var testQueriesTX *Queries
var testDB *pgx.Conn
var testTX *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	testDB, err = pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testTX, err = pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueriesDB = New(testDB)
	testQueriesTX = New(testTX)

	os.Exit(m.Run())
}
