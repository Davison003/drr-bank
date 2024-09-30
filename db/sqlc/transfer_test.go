package db

import (
	"context"
	"testing"
	"time"

	"github.com/Davison003/drr-bank/util"
	"github.com/stretchr/testify/require"
)

var (
	acc1, acc2 Account
)

func createRandomTransfer(t *testing.T, accounts ...Account) Transfer {

	if len(accounts) == 0 {
		acc1 = createRandomAccount(t)
		acc2 = createRandomAccount(t)
	} else {
		acc1 = accounts[0]
		acc2 = accounts[1]
	}

	transferAmount := util.RandomMoney()

	arg := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        transferAmount,
	}

	// creating and testing transfer obj
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	require.Equal(t, acc1.ID, transfer.FromAccountID)
	require.Equal(t, acc2.ID, transfer.ToAccountID)

	// UNNEEDED CODE

	// //withdrawing from FromAccount
	// argFrom := UpdateAccountParams{
	// 	ID: acc1.ID,
	// 	Balance: acc1.Balance - transferAmount,
	// }

	// // receiving amount to ToAccount
	// argTo := UpdateAccountParams{
	// 	ID: acc2.ID,
	// 	Balance: acc2.Balance + transferAmount,
	// }
	// _, errFromAcc := testQueries.UpdateAccount(context.Background(), argFrom)
	// _, errToAcc := testQueries.UpdateAccount(context.Background(), argTo)

	// require.NoError(t, errFromAcc)
	// require.NoError(t, errToAcc)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	// acc1 := createRandomAccount(t)
	// acc2 := createRandomAccount(t)

	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestListTranfers(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, acc1, acc2)
	}

	arg := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, tranf := range transfers {
		require.NotEmpty(t, tranf)
	}
}
