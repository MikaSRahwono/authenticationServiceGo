package usecases

import (
	"authenticationService/internal/domain"
	"authenticationService/pkg/utils"
	"time"
)

type AuthUseCase struct {
	UserRepo      domain.UserRepository
	TokenRepo     domain.TokenRepository
	EmailRepo     domain.EmailRepository
	JwtSecretKey  string
	JwtExpiryTime time.Duration
}

func (a *AuthUseCase) Register(email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := domain.User{
		Email:        email,
		PasswordHash: hashedPassword,
		IsVerified:   false,
		CreatedAt:    time.Now(),
	}

	if err := a.UserRepo.CreateUser(user); err != nil {
		return err
	}

	token, err := utils.GenerateToken(email, a.JwtSecretKey, a.JwtExpiryTime)
	if err != nil {
		return err
	}

	verificationLink := "http://example.com/verify?token=" + token
	return a.EmailRepo.SendEmail(email, "Verify Your Account", verificationLink)
}
