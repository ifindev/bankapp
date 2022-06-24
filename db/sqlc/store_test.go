package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTX(t *testing.T) {
	store := NewStore(testDB)

	// create new accounts
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// parameters to run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	// errs & results are channels to receive errors and transfer result
	// from each transfer goroutine. With these channels, we can use testify
	// to later verify the test result
	errs := make(chan error)
	results := make(chan TransferTxresult)

	// Start running n concurrent transfer transactions with goroutine.
	// We use closure to run the goroutine.
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer

		require.NotEmpty(t, transfer)

		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// GetTransfer with the transfer's ID. Should return no error if transfer exists
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check new fromEntry entry
		fromEntry := result.FromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// check new toEntry entry
		toEntry := result.ToEntry

		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// TODO: check account's balance
	}
}
