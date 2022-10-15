package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trancongthanh1992/samplebank/util"
)

func CreateTransfer(t *testing.T) Transfer {
	accountFrom := CreateRandomAccount(t)
	accountTo := CreateRandomAccount(t)
	// fmt.Println("accountFrom", accountFrom)
	// fmt.Println("accountTo", accountTo)

	arg := CreateTransferParams{
		accountFrom.ID,
		accountTo.ID,
		util.RandomInt(0, accountFrom.Balance),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotEmpty(t, accountFrom.ID)
	require.NotEmpty(t, accountTo.ID)

	require.NotEqual(t, accountFrom.ID, accountTo.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateTransfer(t)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := CreateTransfer(t)

	arg := UpdateTransferParams{
		Amount: transfer.Amount,
		ID:     transfer.ID,
	}
	transfer1, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, transfer.ID, transfer1.ID)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer.Amount, transfer1.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer1.ID)
	require.NotZero(t, transfer1.CreatedAt)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := CreateTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateTransfer(t)

	transfer1, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.Equal(t, transfer.ID, transfer1.ID)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer.Amount, transfer1.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer1.ID)
	require.NotZero(t, transfer1.CreatedAt)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateTransfer(t)
	}
	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
