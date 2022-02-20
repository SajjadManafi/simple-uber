package models

import (
	"time"
)

// Trip is a model for the trips.
type Trip struct {
	ID           int32     `json:"id"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	RiderID      int32     `json:"rider_id"`
	DriverID     int32     `json:"driver_id"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Status       int32     `json:"status"`
	Amount       int64     `json:"amount"`
	CabID        int32     `json:"cab_id"`
	DriverRating int32     `json:"driver_rating"`
}

// CreateTripParams is the input for the CreateTrip function.
type CreateTripParams struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	RiderID     int32  `json:"rider_id"`
	Amount      int64  `json:"amount"`
}

// ListTripsParams is the input for the ListTrips function.
type ListTripsParams struct {
	DriverID int32 `json:"driver_id"`
	RiderID  int32 `json:"rider_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

// UpdateTripDriverAccpetParams is the input for the UpdateTripDriverAccpet function.
type UpdateTripDriverAccpetParams struct {
	ID       int32 `json:"id"`
	DriverID int32 `json:"driver_id"`
	CabID    int32 `json:"cab_id"`
}

// UpdateTripDoneParams is the input for the UpdateTripDone function.
type UpdateTripDoneParams struct {
	ID           int32 `json:"id"`
	DriverRating int32 `json:"driver_rating"`
}

// UpdateTripStatusParams is the input for the UpdateTripStatus function.
type UpdateTripStatusParams struct {
	ID     int32 `json:"id"`
	Status int32 `json:"status"`
}

// TripDoneTransactionParams contains the input parameters of the TripDoneTx function.
type TripDoneTransactionParams struct {
	ID           int32 `json:"id"`
	DriverID     int32 `json:"driver_id"`
	RiderID      int32 `json:"rider_id"`
	Amount       int64 `json:"amount"`
	DriverRating int32 `json:"driver_rating"`
}

// TripDoneTransactionResult contains the result of the TripDoneTx function.
type TripDoneTransactionResult struct {
	Trip   Trip   `json:"trip"`
	User   User   `json:"user"`
	Driver Driver `json:"driver"`
}
