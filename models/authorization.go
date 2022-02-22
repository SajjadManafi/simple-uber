package models

import "time"

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Type     string `json:"type" binding:"required"`
}

type LoginResponse struct {
	AccessToken string            `json:"access_token"`
	User        LoginUserResponse `json:"user"`
}

type LoginUserResponse struct {
	Username string    `json:"username"`
	FullName string    `json:"full_name"`
	Gender   Gender    `json:"gender"`
	Balance  int64     `json:"balance"`
	Email    string    `json:"email"`
	JoinedAt time.Time `json:"joined_at"`
}
