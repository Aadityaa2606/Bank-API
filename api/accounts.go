// This includes functionality for account management operations like creation,
// retrieval, updating, and deletion of bank accounts.
package api

import (
	"errors"
	"net/http"

	db "github.com/Aadityaa2606/simpleBank/db/sqlc"
	"github.com/Aadityaa2606/simpleBank/token"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

// createAccount creates a new bank account for the authenticated user.
// It requires a valid currency in the request body and sets the initial balance to 0.
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

type getAccountByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getAccountById retrieves a specific account by its ID.
// Only returns the account if it belongs to the authenticated user.
func (server *Server) getAccountById(ctx *gin.Context) {
	var req getAccountByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		ctx.JSON(
			http.StatusUnauthorized,
			errorResponse(errors.New("account doesn't belong to the authenticated user")),
		)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	Limit  int32 `form:"limit" binding:"required"`
	Offset int32 `form:"offset"`
}

// getAccounts returns a paginated list of accounts belonging to the authenticated user.
func (server *Server) getAccounts(ctx *gin.Context) {
	var req getAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.Limit,
		Offset: req.Offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

type updateAccountIDRequest struct {
	ID int64 `uri:"id" binding:"required"` // Bind ID from URI
}

type updateAccountBalanceRequest struct {
	Balance int64 `json:"balance" binding:"required"` // Bind balance from JSON body
}

// updateAccount updates the balance of a specific account.
// Only allows updates if the account belongs to the authenticated user.
func (server *Server) updateAccount(ctx *gin.Context) {
	var req1 updateAccountIDRequest
	var req2 updateAccountBalanceRequest

	// Bind URI param (ID)
	if err := ctx.ShouldBindUri(&req1); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Bind JSON body (balance)
	if err := ctx.ShouldBindJSON(&req2); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check if the account belongs to the authenticated user
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	isAccountOwner, err := server.store.CheckAccountOwnership(ctx, db.CheckAccountOwnershipParams{
		ID:    req1.ID,
		Owner: authPayload.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !isAccountOwner {
		ctx.JSON(
			http.StatusUnauthorized,
			errorResponse(errors.New("account doesn't belong to the authenticated user")),
		)
		return
	}

	account, err := server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:      req1.ID,
		Balance: req2.Balance,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type deleteAccountRequest struct {
	ID int64 `uri:"id"`
}

// deleteAccount removes an account from the system.
// Only allows deletion if the account belongs to the authenticated user.
func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check if the account belongs to the authenticated user
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	isAccountOwner, err := server.store.CheckAccountOwnership(ctx, db.CheckAccountOwnershipParams{
		ID:    req.ID,
		Owner: authPayload.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !isAccountOwner {
		ctx.JSON(
			http.StatusUnauthorized,
			errorResponse(errors.New("account doesn't belong to the authenticated user")),
		)
		return
	}

	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
