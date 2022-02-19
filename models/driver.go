package models

import "time"

// Driver is a model for the drivers.
type Driver struct {
	ID             int32     `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FullName       string    `json:"full_name"`
	Gender         Gender    `json:"gender"`
	Balance        int64     `json:"balance"`
	Email          string    `json:"email"`
	CurrentCabID   int32     `json:"current_cab_id"`
	JoinedAt       time.Time `json:"joined_at"`
}

// CreateDriverParams used in creating driver in database
type CreateDriverParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Gender         Gender `json:"gender"`
	Email          string `json:"email"`
}

// ListDriversParams used in get list of drivers from database
type ListDriversParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// UpdateDriverBalanceParams used in update driver balance using driver id
type UpdateDriverBalanceParams struct {
	ID      int32 `json:"id"`
	Balance int64 `json:"balance"`
}

// UpdateDriverCurrentCabParams used in update driver current cab using driver id
type UpdateDriverCurrentCabParams struct {
	ID           int32 `json:"id"`
	CurrentCabID int32 `json:"current_cab_id"`
}

// AddDriverBalanceParams used in add driver balance using driver id
type AddDriverBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int32 `json:"id"`
}
