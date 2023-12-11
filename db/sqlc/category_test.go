package db

import (
	"context"
	"testing"

	"github.com/solracnet/go-finance/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		UserID:      createRandomUser(t).ID,
		Title:       util.RandomString(6),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)
	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.UserID, category2.UserID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestGetCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.UserID, category2.UserID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}
