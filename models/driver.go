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
