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
	dbSource = "postgres://root:1805@localhost:5432/drr_bank?sslmode=disable"
)

var testQueries *Queries
var dbTest *sql.DB

func TestMain(m *testing.M) {
	var err error
	dbTest, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(dbTest)

	os.Exit(m.Run())
}
