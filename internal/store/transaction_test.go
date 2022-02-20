package store

import (
	"context"
	"testing"

	"github.com/SajjadManafi/simple-uber/internal/util"
	"github.com/SajjadManafi/simple-uber/models"
	"github.com/stretchr/testify/require"
)

func TestTripDoneTx(t *testing.T) {
	var err error
	driver := createRandomDriver(t)
	rider := createRandomUser(t)
	cab := createRandomCab(t)

	// test configuration
	rider, err = TestDB.AddUserBalance(context.Background(), models.AddUserBalanceParams{
		ID:     rider.ID,
		Amount: 20000,
	})

	require.NoError(t, err)
	require.NotEmpty(t, rider)

	driver, err = TestDB.UpdateDriverCurrentCab(context.Background(), models.UpdateDriverCurrentCabParams{
		ID:           driver.ID,
		CurrentCabID: cab.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, driver)

	amount := 10

	trip, err := TestDB.CreateTrip(context.Background(), models.CreateTripParams{
		Origin:      util.RandomString(6),
		Destination: util.RandomString(6),
		RiderID:     rider.ID,
		Amount:      int64(amount),
	})

	require.NoError(t, err)
	require.NotEmpty(t, trip)

	trip, err = TestDB.UpdateTripDriverAccpet(context.Background(), models.UpdateTripDriverAccpetParams{
		ID:       trip.ID,
		DriverID: driver.ID,
		CabID:    driver.CurrentCabID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, trip)

	randomRating := util.RandomDriverRating()

	tripResult, err := TestDB.TripDoneTx(context.Background(), models.TripDoneTransactionParams{
		Trip:         trip,
		DriverRating: randomRating,
	})

	require.NoError(t, err)
	require.NotEmpty(t, tripResult)

	// rider
	require.Equal(t, tripResult.User.ID, rider.ID)
	require.Equal(t, tripResult.User.Balance, rider.Balance-trip.Amount)
	// driver
	require.Equal(t, tripResult.Driver.ID, driver.ID)
	require.Equal(t, tripResult.Driver.Balance, driver.Balance+trip.Amount)
	// trip
	require.Equal(t, tripResult.Trip.ID, trip.ID)
	require.Equal(t, tripResult.Trip.Status, int32(4))
	require.Equal(t, tripResult.Trip.DriverRating, randomRating)

}
