package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	db "github.com/Aadityaa2606/Bank-API/db/sqlc"
	"github.com/Aadityaa2606/Bank-API/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store                *db.Store
	tokenMaker           token.Maker
	router               *gin.Engine
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/users/logout", server.logoutUser)
	authRoutes.POST("/users/token/refresh", server.renewAccessToken)
	authRoutes.POST("/users/revoke", server.revokeSession)

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccountById)

	authRoutes.GET("/accounts", server.getAccounts)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.PATCH("/accounts/:id", server.updateAccount)

	authRoutes.POST("/transfer", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
