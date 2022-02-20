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

// CreateUserParams used in creating user in database
type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Gender         Gender `json:"gender"`
	Email          string `json:"email"`
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

// CreateUserRequest used in get request for creating user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Gender   Gender `json:"gender" binding:"required,oneof=male female"`
	Email    string `json:"email" binding:"required,email"`
}

// CreateUserResponse is Response to user creation used in api
type CreateUserResponse struct {
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Gender   Gender    `json:"gender"`
	Balance  int64     `json:"balance"`
	Email    string    `json:"email"`
	JoinedAt time.Time `json:"joined_at"`
}

// GetUserRequest used in get request for getting user
type GetUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type AddUserBalanceRequest struct {
	Amount int64 `json:"amount" binding:"required,min=1"`
}

type DeleteUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}
