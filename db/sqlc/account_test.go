package db

import (
	"context"
	"testing"
	"time"

	"github.com/kazuki-sep27/simple_bank_go/util"
	"github.com/stretchr/testify/require"
)

//CreateRandomAccount generate a radom account and insert to database
func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountsParams {
		Owner    : util.RandomOwner(),
		Balance  : util.RandomBalance(),
		Currency : util.RandomCurrency(),
	}

	result,err := testQuery.CreateAccounts(context.Background(),arg)
	account,accErr := testQuery.GetLastAccount(context.Background())

	require.NoError(t,err)
	require.NoError(t,accErr)
	require.NotEmpty(t,result)

	require.Equal(t, arg.Owner,account.Owner)
	require.Equal(t, arg.Balance,account.Balance)
	require.Equal(t, arg.Currency,account.Currency)

	require.NotZero(t, account.ID)

	return account
}

func TestCreateAccount(t *testing.T){
	CreateRandomAccount(t)
}

func TestGetAccountByID(t *testing.T) {
	
	newAccount := CreateRandomAccount(t)
	checkAccount,err  := testQuery.GetAccountByID(context.Background(),newAccount.ID)

	require.NoError(t,err)
	require.NotEmpty(t,checkAccount)

	require.Equal(t, newAccount.ID,checkAccount.ID)
	require.Equal(t, newAccount.Owner,checkAccount.Owner)
	require.Equal(t, newAccount.Balance,checkAccount.Balance)
	require.Equal(t, newAccount.Currency,checkAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt,checkAccount.CreatedAt, time.Second)
}