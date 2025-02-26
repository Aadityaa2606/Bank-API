package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/Aadityaa2606/simpleBank/db/sqlc"
	"github.com/Aadityaa2606/simpleBank/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string             `json:"username"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusConflict, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)

	ctx.JSON(http.StatusCreated, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	SessionID             string       `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	RefreshToken          string       `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	// Check if the request is valid and bind the request to the req variable
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the user from the database
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Check if the password is correct
	err = util.CheckPassword(user.HashedPassword, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	// Create the access token
	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.accessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create the refresh token
	refreshToken, err := server.tokenMaker.CreateToken(user.Username, server.refreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Get the payload from the access token
	accessPayload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Get the payload from the refresh token
	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("unable to create session: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt.Time,
		User:                  newUserResponse(user),
	})
}

type logoutUserRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (server *Server) logoutUser(ctx *gin.Context) {
	var req logoutUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshTokenPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshTokenPayload.ID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteSession(ctx, session.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "logged out",
	})
}

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest

	// Check if the request is valid and bind the request to the req variable
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Verify the refresh token
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Get the session from the database
	session, err := server.store.GetSession(ctx, payload.ID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("unable to get session: %v", err),
		})
		return
	}

	// Check if the session is revoked
	if session.IsRevoked {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "session has been revoked",
		})
		return
	}

	// Check if the username in the session is the same as the username in the payload
	if session.Username != payload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Create a new access token
	accessToken, err := server.tokenMaker.CreateToken(session.Username, server.accessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: payload.ExpiresAt.Time,
	})
}

type revokeSessionRequest struct {
	SessionID string `json:"session_id" binding:"required"`
}

func (server *Server) revokeSession(ctx *gin.Context) {
	var req revokeSessionRequest

	// Check if the request is valid and bind the request to the req variable
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the session from the database
	session, err := server.store.GetSession(ctx, req.SessionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Check if the session is revoked
	if session.IsRevoked {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "session already revoked",
		})
		return
	}

	// Revoke the session
	err = server.store.RevokeSession(ctx, req.SessionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "session revoked",
	})
}
