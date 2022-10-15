package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trancongthanh1992/samplebank/util"
)

func CreateEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)

	arg := CreateEntryParams{
		account.ID,
		util.RandomInt(0, account.Balance),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateEntry(t)
}

func TestDeleteEntry(t *testing.T) {
	entry := CreateEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
}

func TestUpdateEntry(t *testing.T) {
	entry := CreateEntry(t)
	arg := UpdateEntryParams{
		Amount: entry.Amount,
		ID:     entry.ID,
	}
	entry1, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, entry1.ID, entry.ID)
	require.Equal(t, entry1.Amount, entry.Amount)
	require.Equal(t, entry1.AccountID, entry.AccountID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	entry := CreateEntry(t)

	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.Equal(t, entry1.ID, entry.ID)
	require.Equal(t, entry1.Amount, entry.Amount)
	require.Equal(t, entry1.AccountID, entry.AccountID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry1.ID)
	require.NotZero(t, entry1.CreatedAt)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
