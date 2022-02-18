package models

import "time"

// Cab is a model for the cabs.
type Cab struct {
	ID        int32     `json:"id"`
	DriverID  int32     `json:"driver_id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Color     string    `json:"color"`
	Plate     string    `json:"plate"`
	CreatedAt time.Time `json:"created_at"`
}
