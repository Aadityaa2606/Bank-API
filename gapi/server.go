package gapi

import (
	"fmt"
	"os"
	"time"

	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/pb"
	"github.com/Aadityaa2606/Bank-API/token"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	store                *db.Store
	tokenMaker           token.Maker
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	accessTokenDuration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		return nil, fmt.Errorf("cannot parse access token duration: %w", err)
	}

	refreshTokenDuration, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
	if err != nil {
		return nil, fmt.Errorf("cannot parse refresh token duration: %w", err)
	}

	server := &Server{
		store:                store,
		tokenMaker:           tokenMaker,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}

	return server, nil
}
