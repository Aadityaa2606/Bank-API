package gapi

import (
	"context"
	"time"

	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/pb"
	"github.com/Aadityaa2606/Bank-API/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username cannot be empty")
	}

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %s", err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized: cannot update other users")
	}

	violations := validateUpdateUserRequest(req)

	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateUserParams{
		Username: req.Username,
		FullName: pgtype.Text{
			String: req.GetFullName(),
			Valid:  req.GetFullName() != "",
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.GetEmail() != "",
		},
	}

	if req.GetPassword() != "" {
		hashedPassword, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Unimplemented, "failed to hash password: %s", err)
		}
		arg.HashedPassword = pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		}
		arg.PasswordChangedAt = pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		}

	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
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

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolations("username", err))
	}

	if req.GetPassword() != "" {
		if err := util.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolations("password", err))
		}
	}

	if req.GetFullName() != "" {
		if err := util.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolations("full_name", err))
		}
	}

	if req.GetEmail() != "" {
		if err := util.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolations("email", err))
		}
	}

	return violations
}
