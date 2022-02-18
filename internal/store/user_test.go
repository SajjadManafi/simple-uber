package store

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/SajjadManafi/simple-uber/models"
	"github.com/stretchr/testify/require"
)

// createRandomUser creates random user in database
func createRandomUser(t *testing.T) models.User {
	HashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)
	require.NotEmpty(t, HashedPassword)

	args := models.CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: HashedPassword,
		FullName:       util.RandomString(8),
		Gender:         util.RandomGender(),
		Email:          util.RandomEmail(),
	}

	user, err := TestDB.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Gender, user.Gender)
	require.Equal(t, args.Email, user.Email)

	require.NotZero(t, user.ID)
	require.Zero(t, user.Balance)
	require.NotZero(t, user.JoinedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	// create random user
	user1 := createRandomUser(t)

	// get user
	user2, err := TestDB.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.Balance, user2.Balance)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.JoinedAt, user2.JoinedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	// create random user
	createRandomUser(t)

	// list users
	users, err := TestDB.ListUsers(context.Background(), models.ListUsersParams{
		Limit:  1,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, users)

	require.Len(t, users, 1)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	args := models.UpdateUserParams{
		ID:      user1.ID,
		Balance: util.RandomInt(1000, 2000),
	}

	user2, err := TestDB.UpdateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, args.ID, user2.ID)
	require.Equal(t, args.Balance, user2.Balance)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := TestDB.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := TestDB.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.Empty(t, user2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

// TODO negative balance tests
func TestAddUserBalance(t *testing.T) {
	user1 := createRandomUser(t)

	args := models.AddUserBalanceParams{
		ID:     user1.ID,
		Amount: util.RandomInt(1000, 2000),
	}

	user2, err := TestDB.AddUserBalance(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, args.ID, user2.ID)
	require.Equal(t, user1.Balance+args.Amount, user2.Balance)
}
