package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user *entities.User) error {
	err := r.db.Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
