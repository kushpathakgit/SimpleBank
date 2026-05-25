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
	dbSource = "postgresql://root:mysecret@127.0.0.1:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB // <-- Make sure this exists!

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)
	testDB = conn // <-- Add this exact line! This was missing.

	os.Exit(m.Run())
}
