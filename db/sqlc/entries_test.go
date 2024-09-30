package db

import (
	"context"
	"testing"
	"time"

	"github.com/Davison003/drr-bank/util"
	"github.com/stretchr/testify/require"
)

var acc Account

func createRandomEntry(t *testing.T, accounts ...Account) Entry {
	if len(accounts) == 0 {
		acc = createRandomAccount(t)
	} else {
		acc = accounts[0]
	}

	randAmount := util.RandomMoney()

	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    randAmount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	argAcc := UpdateAccountParams{
		ID:      acc.ID,
		Balance: acc.Balance + randAmount,
	}

	_, errAcc := testQueries.UpdateAccount(context.Background(), argAcc)
	require.NoError(t, errAcc)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	accFromTestList := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, accFromTestList)
	}

	arg := ListEntriesParams{
		AccountID: accFromTestList.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, e := range entries {
		require.NotEmpty(t, e)
	}

}
