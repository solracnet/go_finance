package db

import (
	"context"
	"testing"
	"time"

	"github.com/solracnet/go-finance/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	category := createRandomCategory(t)
	arg := CreateAccountParams{
		UserID:      category.UserID,
		CategoryID:  category.ID,
		Title:       util.RandomString(6),
		Type:        "debit",
		Description: util.RandomString(20),
		Value:       20,
		Date:        time.Now(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.CategoryID, account.CategoryID)
	require.Equal(t, arg.Title, account.Title)
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Description, account.Description)
	require.Equal(t, arg.Value, account.Value)
	require.NotEmpty(t, account.Date)
	require.NotEmpty(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountById(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Description, account2.Description)
	require.Equal(t, account1.Value, account2.Value)
	require.Equal(t, account1.Date, account2.Date)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestGetAccountById(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountById(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Description, account2.Description)
	require.Equal(t, account1.Value, account2.Value)
	require.Equal(t, account1.Date, account2.Date)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:          account1.ID,
		Title:       util.RandomString(6),
		Description: util.RandomString(20),
		Value:       20,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.NotEqual(t, account1.Title, account2.Title)
	require.NotEqual(t, account1.Description, account2.Description)
	require.Equal(t, account1.Value, account2.Value)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 5; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := GetAccountsParams{
		UserID:      lastAccount.UserID,
		CategoryID:  lastAccount.CategoryID,
		Type:        lastAccount.Type,
		Title:       lastAccount.Title,
		Description: lastAccount.Description,
		Date:        lastAccount.Date,
	}

	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.GreaterOrEqual(t, len(accounts), 1)
	for _, category := range accounts {
		require.Equal(t, lastAccount.ID, category.ID)
	}
}
