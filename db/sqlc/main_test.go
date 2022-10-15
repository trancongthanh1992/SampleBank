package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/trancongthanh1992/samplebank/util"
)

var testQueries *Queries
var testDb *sql.DB

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:aA@123123@localhost:5432/sample_bank_db?sslmode=disable"
// )

// Entry point
func TestMain(m *testing.M) {

	// multiple transaction raise error because `:= declare new variable in block scope`.
	// global scope on file `store_test.go` doesn't effect.
	// testDb, err := sql.Open(dbDriver, dbSource)

	// always redeclare variable each connection

	var config, err = util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
