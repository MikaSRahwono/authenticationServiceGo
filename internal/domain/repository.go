package domain

import (
	"authenticationService/internal/domain/entities"
	"time"
)

type UserRepository interface {
	CreateUser(user entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user entities.User) error
}

type TokenRepository interface {
	StoreToken(token string, expiration time.Time) error
	ValidateToken(token string) (bool, error)
}

type EmailRepository interface {
	SendEmail(to, subject, body string) error
}
