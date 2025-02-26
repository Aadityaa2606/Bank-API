package main

import (
	"context"
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
	dbSource := os.Getenv("DB_SOURCE")

	// Creates DB connection
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store)

	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(os.Getenv("SERVER_ADDR"))
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
