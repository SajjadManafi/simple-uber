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
