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

// CreateCabParams is a model for the parameters of creating a cab.
type CreateCabParams struct {
	DriverID int32  `json:"driver_id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Color    string `json:"color"`
	Plate    string `json:"plate"`
}

// ListCabsParams is a model for the parameters of listing cabs.
type ListCabsParams struct {
	DriverID int32 `json:"driver_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

// UpdateCabsDriverIdParams is a model for the parameters of updating the driver id of a cab.
type UpdateCabsDriverIdParams struct {
	ID       int32 `json:"id"`
	DriverID int32 `json:"driver_id"`
}
