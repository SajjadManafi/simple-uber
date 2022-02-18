package models

import "time"

// User is model for users
type User struct {
	ID             int32  `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Gender         Gender `json:"gender"`
	// must be positive
	Balance  int64     `json:"balance"`
	Email    string    `json:"email"`
	JoinedAt time.Time `json:"joined_at"`
}
