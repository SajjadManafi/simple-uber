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

// createRandomDriver creates random driver in database
func createRandomDriver(t *testing.T) models.Driver {
	HashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)
	require.NotEmpty(t, HashedPassword)

	args := models.CreateDriverParams{
		Username:       util.RandomUsername(),
		HashedPassword: HashedPassword,
		FullName:       util.RandomString(8),
		Gender:         util.RandomGender(),
		Email:          util.RandomEmail(),
	}

	driver, err := TestDB.CreateDriver(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, driver)

	require.Equal(t, args.Username, driver.Username)
	require.Equal(t, args.HashedPassword, driver.HashedPassword)
	require.Equal(t, args.FullName, driver.FullName)
	require.Equal(t, args.Gender, driver.Gender)
	require.Equal(t, args.Email, driver.Email)

	require.NotZero(t, driver.ID)
	require.Zero(t, driver.Balance)
	require.Zero(t, driver.CurrentCabID)
	require.NotZero(t, driver.JoinedAt)

	return driver
}

func TestCreateDriver(t *testing.T) {
	createRandomDriver(t)
}

func TestGetDriver(t *testing.T) {
	// create random driver
	driver1 := createRandomDriver(t)

	// get user
	driver2, err := TestDB.GetDriver(context.Background(), driver1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, driver2)

	require.Equal(t, driver1.ID, driver2.ID)
	require.Equal(t, driver1.Username, driver2.Username)
	require.Equal(t, driver1.HashedPassword, driver2.HashedPassword)
	require.Equal(t, driver1.FullName, driver2.FullName)
	require.Equal(t, driver1.Gender, driver2.Gender)
	require.Equal(t, driver1.Balance, driver2.Balance)
	require.Equal(t, driver1.Email, driver2.Email)
	require.Equal(t, driver1.CurrentCabID, driver2.CurrentCabID)
	require.WithinDuration(t, driver1.JoinedAt, driver2.JoinedAt, time.Second)
}

func TestListDrivers(t *testing.T) {
	// create random driver
	createRandomDriver(t)

	// list drivers
	drivers, err := TestDB.ListDrivers(context.Background(), models.ListDriversParams{
		Limit:  1,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, drivers)

	require.Len(t, drivers, 1)
}

func TestUpdateDriverBalance(t *testing.T) {
	// create random driver
	driver1 := createRandomDriver(t)

	args := models.UpdateDriverBalanceParams{
		ID:      driver1.ID,
		Balance: util.RandomInt(100, 1000),
	}
	// update balance
	driver2, err := TestDB.UpdateDriverBalance(context.Background(), args)
	require.NoError(t, err)

	require.NoError(t, err)
	require.NotEmpty(t, driver2)

	require.Equal(t, args.ID, driver2.ID)
	require.Equal(t, args.Balance, driver2.Balance)
}

func TestUpdateDriverCurrentCab(t *testing.T) {
	// create random driver
	driver1 := createRandomDriver(t)

	// create random cab
	cab := createRandomCab(t)

	// update cab driver id
	cab2, err := TestDB.UpdateCabsDriverId(context.Background(), models.UpdateCabsDriverIdParams{
		ID:       cab.ID,
		DriverID: driver1.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, cab2)

	// update driver current cab id
	driver2, err := TestDB.UpdateDriverCurrentCab(context.Background(), models.UpdateDriverCurrentCabParams{
		ID:           driver1.ID,
		CurrentCabID: cab.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, driver2)

	require.Equal(t, driver1.ID, driver2.ID)
	require.Equal(t, cab.ID, driver2.CurrentCabID)
}

func TestAddDriverBalance(t *testing.T) {
	driver1 := createRandomDriver(t)

	args := models.AddDriverBalanceParams{
		ID:     driver1.ID,
		Amount: util.RandomInt(1000, 2000),
	}

	user2, err := TestDB.AddDriverBalance(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, args.ID, user2.ID)
	require.Equal(t, driver1.Balance+args.Amount, user2.Balance)
}

func TestDeleteDriver(t *testing.T) {
	driver1 := createRandomDriver(t)

	err := TestDB.DeleteDriver(context.Background(), driver1.ID)
	require.NoError(t, err)

	driver2, err := TestDB.GetDriver(context.Background(), driver1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, driver2)
}
