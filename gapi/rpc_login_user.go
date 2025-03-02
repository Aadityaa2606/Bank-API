package gapi

import (
	"context"

	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/pb"
	"github.com/Aadityaa2606/Bank-API/util"
	"github.com/Aadityaa2606/Bank-API/val"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginUserRequest(req)

	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "cannot get user: %v", err)
	}

	err = util.CheckPassword(user.HashedPassword, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Create the access token
	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.accessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create access token: %v", err)
	}

	// Create the refresh token
	refreshToken, err := server.tokenMaker.CreateToken(user.Username, server.refreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create refresh token: %v", err)
	}

	// Get the payload from the access token
	accessPayload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot verify access token: %v", err)
	}

	// Get the payload from the refresh token
	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot verify refresh token: %v", err)
	}

	// Create a new session in the database
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.RegisteredClaims.ID,
		Username:     refreshPayload.Username,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    pgtype.Timestamptz{Time: refreshPayload.RegisteredClaims.ExpiresAt.Time, Valid: true},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create session: %v", err)
	}

	return &pb.LoginUserResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		SessionId:             session.ID,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.RegisteredClaims.ExpiresAt.Time),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.RegisteredClaims.ExpiresAt.Time),
		User: &pb.User{
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordChangedAt.Time),
			CreatedAt:         timestamppb.New(user.CreatedAt.Time),
		},
	}, nil
}

func validateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolations("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolations("password", err))
	}

	return violations
}
