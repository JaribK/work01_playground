package auth

import (
	"work01/internal/entities"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	GetUserByEmail(email string) (*entities.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
