package db

import (
	"context"
	"testing"

	"github.com/kazuki-sep27/simple_bank_go/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T){
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
}