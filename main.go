package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/Aadityaa2606/Bank-API/api"
	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/gapi"
	"github.com/Aadityaa2606/Bank-API/pb"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	if os.Getenv("SERVER_MODE") == "http" {
		runGinServer(store)
	} else {
		runGrpcServer(store)
	}
}

func runGrpcServer(store *db.Store) {
	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", os.Getenv("GRPC_SERVER_ADDR"))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	log.Printf("starting gRPC server on %s", os.Getenv("GRPC_SERVER_ADDR"))
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

func runGinServer(store *db.Store) {
	server, err := api.NewServer(store)

	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	log.Printf("starting HTTP/REST server on %s", os.Getenv("HTTP_SERVER_ADDR"))
	err = server.Start(os.Getenv("HTTP_SERVER_ADDR"))
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
