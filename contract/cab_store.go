package contract

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

type CabStore interface {
	CreateCab(ctx context.Context, arg models.CreateCabParams) (models.Cab, error)
	GetCab(ctx context.Context, id int32) (models.Cab, error)
	ListCabs(ctx context.Context, arg models.ListCabsParams) ([]models.Cab, error)
	UpdateCabsDriverId(ctx context.Context, arg models.UpdateCabsDriverIdParams) (models.Cab, error)
	DeleteCab(ctx context.Context, id int32) error
}
