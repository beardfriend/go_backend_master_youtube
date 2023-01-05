package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5434/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDb      *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
