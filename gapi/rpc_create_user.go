package gapi

import (
	"context"

	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/pb"
	"github.com/Aadityaa2606/Bank-API/util"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserResquest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {

		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505": // unique_violation
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}

		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: &pb.User{
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordChangedAt.Time),
			CreatedAt:         timestamppb.New(user.CreatedAt.Time),
		},
	}

	return rsp, nil
}
