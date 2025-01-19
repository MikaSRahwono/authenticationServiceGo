package entities

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	IsVerified   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
