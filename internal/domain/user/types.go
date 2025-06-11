package user

import "time"

type User struct {
	CreatedAt    time.Time
	ID           string
	Email        string
	PasswordHash string
	RoleID       int
}

type Input struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
