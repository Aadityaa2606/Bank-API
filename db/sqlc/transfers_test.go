package db

import (
	"context"
	"testing"
	"time"

	"github.com/Aadityaa2606/Bank-API/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	require.NotEmpty(t, account2)

	transferArg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), transferArg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transferArg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferArg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferArg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfersByToAccountID(t *testing.T) {
	trasfer1 := createRandomTransfer(t)

	args := GetTransferByToAccountIDParams{
		ToAccountID: trasfer1.ToAccountID,
		Limit:       1000,
	}
	transfers, err := testQueries.GetTransferByToAccountID(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	found := false
	for _, transfer := range transfers {
		if transfer.ID == trasfer1.ID {
			found = true
			require.Equal(t, trasfer1.FromAccountID, transfer.FromAccountID)
			require.Equal(t, trasfer1.ToAccountID, transfer.ToAccountID)
			require.Equal(t, trasfer1.Amount, transfer.Amount)
			break
		}
	}
	require.True(t, found)
}

func TestGetTransfersByFromAccountID(t *testing.T) {
	trasfer1 := createRandomTransfer(t)

	args := GetTransferByFromAccountIDParams{
		FromAccountID: trasfer1.FromAccountID,
		Limit:         1000,
	}
	transfers, err := testQueries.GetTransferByFromAccountID(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	found := false
	for _, transfer := range transfers {
		if transfer.ID == trasfer1.ID {
			found = true
			require.Equal(t, trasfer1.FromAccountID, transfer.FromAccountID)
			require.Equal(t, trasfer1.ToAccountID, transfer.ToAccountID)
			require.Equal(t, trasfer1.Amount, transfer.Amount)
			break
		}
	}
	require.True(t, found)
}

// Test Logic : We are creating a random transfer and checking if listTransfer returns non empty result
func TestListTransfers(t *testing.T) {
	createRandomTransfer(t)

	args := ListTransfersParams{
		Limit: 10,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	args := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, args.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	_, err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransferByID(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
