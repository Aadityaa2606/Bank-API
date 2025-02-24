package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Aadityaa2606/simpleBank/api"
	db "github.com/Aadityaa2606/simpleBank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Loads environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbSource := fmt.Sprintf("postgresql://%v:%v@localhost:5432/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_NAME"))

	// Creates DB connection
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(os.Getenv("SERVER_ADDR"))
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
