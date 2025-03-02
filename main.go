package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Aadityaa2606/Bank-API/api"
	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/gapi"
	"github.com/Aadityaa2606/Bank-API/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// Loads environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}
	dbSource := os.Getenv("DB_SOURCE")

	if os.Getenv("ENVIROMENT") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Creates DB connection
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	store := db.NewStore(conn)

	if os.Getenv("SERVER_MODE") == "http" {
		runGinServer(store)
	} else {
		go runGatewayServer(store)
		runGrpcServer(store)
	}
}

func runGrpcServer(store *db.Store) {
	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server: ")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLogger)

	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", os.Getenv("GRPC_SERVER_ADDR"))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server: ")
	}

	log.Info().Msgf("starting gRPC server on %s", os.Getenv("GRPC_SERVER_ADDR"))
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server: ")
	}
}

func runGatewayServer(store *db.Store) {
	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server: ")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register service: ")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", os.Getenv("HTTP_SERVER_ADDR"))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server: ")
	}

	log.Info().Msgf("starting HTTP gateway server on %s", os.Getenv("HTTP_SERVER_ADDR"))
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP Gateway server: ")
	}
}

func runGinServer(store *db.Store) {
	server, err := api.NewServer(store)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server: ")
	}

	log.Info().Msgf("starting HTTP/REST server on %s", os.Getenv("HTTP_SERVER_ADDR"))
	err = server.Start(os.Getenv("HTTP_SERVER_ADDR"))
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start server: ")
	}
}
