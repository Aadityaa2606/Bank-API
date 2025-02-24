package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbSource := fmt.Sprintf("postgresql://%v:%v@localhost:5432/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_NAME"))

	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
