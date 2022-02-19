package contract

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

type TripStore interface {
	CreateTrip(ctx context.Context, arg models.CreateTripParams) (models.Trip, error)
	GetTrip(ctx context.Context, id int32) (models.Trip, error)
	ListTrips(ctx context.Context, arg models.ListTripsParams) ([]models.Trip, error)
	UpdateTripStatus(ctx context.Context, arg models.UpdateTripStatusParams) (models.Trip, error)
	UpdateTripDriverAccpet(ctx context.Context, arg models.UpdateTripDriverAccpetParams) (models.Trip, error)
	UpdateTripDone(ctx context.Context, arg models.UpdateTripDoneParams) (models.Trip, error)
	DeleteTrip(ctx context.Context, id int32) error
}
