package db

import (
	"context"
	"testing"
	"time"

	"github.com/invokify/go-practice-database/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	arg := createRandomTransfer(t, createRandomAccount(t).ID)

	require.NotEmpty(t, arg)
	require.NotZero(t, arg.ID)
	require.NotZero(t, arg.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t, createRandomAccount(t).ID)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account.ID)
	}

	arg := ListTransfersParams{
		FromAccountID: account.ID,
		ToAccountID:   account.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Equal(t, len(transfers), 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func createRandomTransfer(t *testing.T, accountID int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: accountID,
		ToAccountID:   accountID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
