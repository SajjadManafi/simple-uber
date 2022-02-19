package contract

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

type DriverStore interface {
	CreateDriver(ctx context.Context, arg models.CreateDriverParams) (models.Driver, error)
	GetDriver(ctx context.Context, id int32) (models.Driver, error)
	ListDrivers(ctx context.Context, arg models.ListDriversParams) ([]models.Driver, error)
	UpdateDriverBalance(ctx context.Context, arg models.UpdateDriverBalanceParams) (models.Driver, error)
	UpdateDriverCurrentCab(ctx context.Context, arg models.UpdateDriverCurrentCabParams) (models.Driver, error)
	AddDriverBalance(ctx context.Context, arg models.AddDriverBalanceParams) (models.Driver, error)
	DeleteDriver(ctx context.Context, id int32) error
}
