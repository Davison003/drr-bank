package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Davison003/drr-bank/util"
	_ "github.com/lib/pq"
)



var testQueries *Queries
var dbTest *sql.DB

func TestMain(m *testing.M) {
	
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config files: ", err)
	}

	dbTest, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(dbTest)

	os.Exit(m.Run())
}
