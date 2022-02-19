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

// createRandomCab creates random cab in database
func createRandomCab(t *testing.T) models.Cab {
	driver := createRandomDriver(t)

	args := models.CreateCabParams{
		DriverID: driver.ID,
		Brand:    util.RandomCabBrand(),
		Model:    util.RandomCabModel(),
		Color:    util.RandomCabColor(),
		Plate:    util.RandomCabPlate(),
	}

	cab, err := TestDB.CreateCab(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, driver)

	require.Equal(t, args.DriverID, cab.DriverID)
	require.Equal(t, args.Brand, cab.Brand)
	require.Equal(t, args.Model, cab.Model)
	require.Equal(t, args.Color, cab.Color)
	require.Equal(t, args.Plate, cab.Plate)

	require.NotZero(t, cab.ID)
	require.NotZero(t, cab.CreatedAt)

	return cab
}

func TestCreateCab(t *testing.T) {
	createRandomCab(t)
}

func TestGetCab(t *testing.T) {
	// create random cab
	cab := createRandomCab(t)

	// get cab
	cab2, err := TestDB.GetCab(context.Background(), cab.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cab2)

	require.Equal(t, cab.ID, cab2.ID)
	require.Equal(t, cab.DriverID, cab2.DriverID)
	require.Equal(t, cab.Brand, cab2.Brand)
	require.Equal(t, cab.Model, cab2.Model)
	require.Equal(t, cab.Color, cab2.Color)
	require.Equal(t, cab.Plate, cab2.Plate)
	require.WithinDuration(t, cab.CreatedAt, cab2.CreatedAt, time.Second)
}

func TestListCabs(t *testing.T) {
	var lastCab models.Cab
	for i := 0; i < 10; i++ {
		lastCab = createRandomCab(t)
	}

	arg := models.ListCabsParams{
		DriverID: lastCab.DriverID,
		Limit:    5,
		Offset:   0,
	}

	cabs, err := TestDB.ListCabs(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, cabs)

	for _, cab := range cabs {
		require.NotEmpty(t, cab)
		require.Equal(t, lastCab.DriverID, cab.DriverID)
	}

}

func TestUpdateCabsDriverId(t *testing.T) {

	cab := createRandomCab(t)

	// update cab

	// create random driver
	driver := createRandomDriver(t)

	arg := models.UpdateCabsDriverIdParams{
		ID:       cab.ID,
		DriverID: driver.ID,
	}
	cab2, err := TestDB.UpdateCabsDriverId(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cab2)

	require.Equal(t, cab.ID, cab2.ID)
	require.Equal(t, arg.DriverID, cab2.DriverID)

}

func TestDeleteCab(t *testing.T) {
	cab := createRandomCab(t)

	err := TestDB.DeleteCab(context.Background(), cab.ID)
	require.NoError(t, err)

	cab2, err := TestDB.GetCab(context.Background(), cab.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, cab2)
}
