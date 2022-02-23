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

// CreateDriverRequest used in get request for creating driver
type CreateDriverRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Gender   Gender `json:"gender" binding:"required,oneof=male female"`
	Email    string `json:"email" binding:"required,email"`
}

// CreateDriverResponse is Response to driver creation used in api
type CreateDriverResponse struct {
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Gender   Gender    `json:"gender"`
	Balance  int64     `json:"balance"`
	Email    string    `json:"email"`
	JoinedAt time.Time `json:"joined_at"`
}

// GetDriverRequest used in get request for getting driver
type GetDriverRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// DriverBalanceWithdrawRequest used in get request for withdrawing driver balance
type DriverBalanceWithdrawRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// DriverBalanceWithdrawResponse is Response to driver balance withdrawal used in api
type DriverBalanceWithdrawResponse struct {
	Username             string `json:"username"`
	FullName             string `json:"full_name"`
	Balance              int64  `json:"balance_before_withdraw"`
	BalanceAfterWithdraw int64  `json:"balance_after_withdraw"`
}

// SetCabRequest used in get request for setting cab for driver
type SetCabRequest struct {
	CabID int32 `json:"cab_id" binding:"required,min=1"`
}
