package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T){
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	fmt.Println(">>Before transaction : ", account1.Balance, account2.Balance)

	//run test n concurrent transfer transaction
	n := 1

	errs := make(chan error)
	results := make(chan TransferTxResult)
	amounts := make(chan int64)

	for i:=0; i<n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(),txKey,txName)

			//amount := util.RandomBalance()
			amount := int64(10)
			
			result,err := store.TransferTx(ctx,CreateTransferParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: amount,
			})

			errs <- err
			results <- result
			amounts <- amount
		}()
	}

	//check result
	sumOfAmount := int64(0)

	for i:=0; i<n; i++ {
		err := <- errs
		require.NoError(t,err)

		result := <- results
		require.NotEmpty(t,result)

		amount := <- amounts

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t,transfer)
		
		require.Equal(t, account1.ID,transfer.FromAccountID)
		require.Equal(t, account2.ID,transfer.ToAccountID)
		require.Equal(t, amount,transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		//check entry
		fromEntry := result.FromEntry
		require.NotEmpty(t,fromEntry)

		require.Equal(t, account1.ID,fromEntry.AccountID)
		require.Equal(t, -amount,fromEntry.Amount)
		require.NotZero(t, fromEntry.AccountID)
		require.NotZero(t, fromEntry.CreatedAt)

		_,err = store.GetEntryByID(context.Background(),fromEntry.ID)
		require.NoError(t,err)

		toEntry := result.ToEntry
		require.NotEmpty(t,toEntry)

		require.Equal(t, account2.ID,toEntry.AccountID)
		require.Equal(t, amount,toEntry.Amount)
		require.NotZero(t, toEntry.AccountID)
		require.NotZero(t, toEntry.CreatedAt)

		_,err = store.GetEntryByID(context.Background(),toEntry.ID)
		require.NoError(t,err)

		//check account
		fromAccount := result.FromAccount
		require.NotEmpty(t,fromAccount)
		require.Equal(t, account1.ID,fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t,toAccount)
		require.Equal(t, account2.ID,toAccount.ID)

		//check account's balance
		fmt.Println(">>Loop transaction : ", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance

		require.Equal(t, diff1,diff2)
		require.True(t,diff1>0)
		require.True(t,diff2>0)
		require.True(t,diff1%amount == 0) //  1 * amount, 2 * amount, ..., n * amount 

		sumOfAmount = sumOfAmount+amount
	}
	
	//check the final update balance	
	fmt.Println(">>After transaction : ", account1.Balance, account2.Balance)

	updateAccount1, err := testQuery.GetAccountByID(context.Background(),account1.ID)
	require.NoError(t,err)

	updateAccount2, err := testQuery.GetAccountByID(context.Background(),account2.ID)
	require.NoError(t,err)

	require.Equal(t, account1.Balance-sumOfAmount,updateAccount1.Balance)
	require.Equal(t, account2.Balance+sumOfAmount,updateAccount2.Balance)

}