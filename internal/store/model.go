package store

import "github.com/SajjadManafi/simple-uber/models"

// CreateUserParams used in creating user in database
type CreateUserParams struct {
	Username       string        `json:"username"`
	HashedPassword string        `json:"hashed_password"`
	FullName       string        `json:"full_name"`
	Gender         models.Gender `json:"gender"`
	Email          string        `json:"email"`
}

// ListUsersParams used in get list of users
type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// UpdateUserParams used in update user balance using user id
type UpdateUserParams struct {
	ID      int32 `json:"id"`
	Balance int64 `json:"balance"`
}

// AddUserBalanceParams used in add user balance using user id
type AddUserBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int32 `json:"id"`
}