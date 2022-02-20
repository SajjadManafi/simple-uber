package contract

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

type Store interface {
	UserStore
	DriverStore
	CabStore
	TripStore
	TripDoneTx(ctx context.Context, arg models.TripDoneTransactionParams) (models.TripDoneTransactionResult, error)
}
