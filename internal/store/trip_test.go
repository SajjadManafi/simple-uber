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

func createRandomTrips(t *testing.T) models.Trip {
	// create random user
	user1 := createRandomUser(t)

	args := models.CreateTripParams{
		Origin:      util.RandomString(8),
		Destination: util.RandomString(8),
		RiderID:     user1.ID,
		Amount:      util.RandomInt(1, 100),
	}

	trip, err := TestDB.CreateTrip(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, trip)

	require.Equal(t, args.Origin, trip.Origin)
	require.Equal(t, args.Destination, trip.Destination)
	require.Equal(t, args.RiderID, trip.RiderID)
	require.Equal(t, args.Amount, trip.Amount)

	require.NotZero(t, trip.ID)
	require.Zero(t, trip.DriverID)
	require.Zero(t, trip.CabID)
	require.Zero(t, trip.DriverRating)
	require.NotZero(t, trip.StartTime)
	require.Zero(t, trip.EndTime)
	require.Equal(t, trip.Status, int32(1))

	return trip

}

func TestCreateTrip(t *testing.T) {
	createRandomTrips(t)
}

func TestGetTrip(t *testing.T) {

	// create random trip
	trip := createRandomTrips(t)

	// get trip
	trip2, err := TestDB.GetTrip(context.Background(), trip.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trip2)

	require.Equal(t, trip.ID, trip2.ID)
	require.Equal(t, trip.Origin, trip2.Origin)
	require.Equal(t, trip.Destination, trip2.Destination)
	require.Equal(t, trip.RiderID, trip2.RiderID)
	require.Equal(t, trip.DriverID, trip2.DriverID)
	require.Equal(t, trip.Amount, trip2.Amount)
	require.Equal(t, trip.DriverRating, trip2.DriverRating)
	require.WithinDuration(t, trip.StartTime, trip2.StartTime, time.Second)
}

func TestListTrips(t *testing.T) {
	var lastTrip models.Trip
	for i := 0; i < 10; i++ {
		lastTrip = createRandomTrips(t)
	}

	arg := models.ListTripsParams{
		RiderID: lastTrip.RiderID,
		Limit:   5,
		Offset:  0,
	}

	trips, err := TestDB.ListTrips(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trips)

	for _, trip := range trips {
		require.NotEmpty(t, trip)
		require.Equal(t, lastTrip.DriverID, trip.DriverID)
	}

}

func TestUpdateTripStatus(t *testing.T) {
	trip := createRandomTrips(t)

	arg := models.UpdateTripStatusParams{
		ID:     trip.ID,
		Status: 2,
	}

	trip2, err := TestDB.UpdateTripStatus(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, trip2)

	require.Equal(t, arg.Status, trip2.Status)
	require.Equal(t, trip.ID, trip2.ID)

}

func TestUpdateTripDriverAccpet(t *testing.T) {
	// create random driver
	driver := createRandomDriver(t)
	// create random cab
	cab := createRandomCab(t)

	driver, err := TestDB.UpdateDriverCurrentCab(context.Background(), models.UpdateDriverCurrentCabParams{
		ID:           driver.ID,
		CurrentCabID: cab.ID,
	})
	require.NoError(t, err)

	trip := createRandomTrips(t)

	arg := models.UpdateTripDriverAccpetParams{
		ID:       trip.ID,
		DriverID: driver.ID,
		CabID:    driver.CurrentCabID,
	}

	trip2, err := TestDB.UpdateTripDriverAccpet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, trip2)

	require.Equal(t, arg.DriverID, trip2.DriverID)
	require.Equal(t, arg.CabID, trip2.CabID)
	require.Equal(t, trip.ID, trip2.ID)
	require.Equal(t, int32(2), trip2.Status)

}

func TestUpdateTripDone(t *testing.T) {
	// create random driver
	driver := createRandomDriver(t)
	// create random cab
	cab := createRandomCab(t)

	driver, err := TestDB.UpdateDriverCurrentCab(context.Background(), models.UpdateDriverCurrentCabParams{
		ID:           driver.ID,
		CurrentCabID: cab.ID,
	})
	require.NoError(t, err)

	trip := createRandomTrips(t)

	arg := models.UpdateTripDoneParams{
		ID:           trip.ID,
		DriverRating: int32(util.RandomInt(1, 5)),
	}

	trip2, err := TestDB.UpdateTripDone(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trip2)

	require.Equal(t, trip.ID, trip2.ID)
	require.Equal(t, int32(4), trip2.Status)
	require.Equal(t, arg.DriverRating, trip2.DriverRating)

}

func TestDeleteTrip(t *testing.T) {
	trip := createRandomTrips(t)

	err := TestDB.DeleteTrip(context.Background(), trip.ID)
	require.NoError(t, err)

	trip2, err := TestDB.GetTrip(context.Background(), trip.ID)
	require.Error(t, err)
	require.Empty(t, trip2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
