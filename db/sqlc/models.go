// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64              `json:"id"`
	Owner     string             `json:"owner"`
	Currency  string             `json:"currency"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Balance   int64              `json:"balance"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be -ve or +ve
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Session struct {
	ID           string             `json:"id"`
	Username     string             `json:"username"`
	RefreshToken string             `json:"refresh_token"`
	IsRevoked    bool               `json:"is_revoked"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	ExpiresAt    pgtype.Timestamptz `json:"expires_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be +ve
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	Username          string             `json:"username"`
	HashedPassword    string             `json:"hashed_password"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}
