package repositories

import (
	"github.com/MikaSRahwono/authenticationServiceGo/internal/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user entities.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := r.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (r *UserRepository) UpdateUser(user entities.User) error {
	return r.DB.Save(&user).Error
}
