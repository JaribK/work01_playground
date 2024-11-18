package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(user *entities.User) error
	GetById(id uuid.UUID) (*entities.User, error)
	GetAll(page, size int) ([]entities.User, int64, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID, deleteBy uuid.UUID) error
	GetUserByEmail(email string) (*entities.User, error)
	IsEmailExists(email string) (bool, error)
	IsPhoneExists(phone string) (bool, error)
	IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error)
	IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = r.db.Create(&user).Error
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

func (r *userRepository) GetAll(page, size int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	offset := (page - 1) * size
	err := r.db.Model(&entities.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Role").Limit(size).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) Update(user *entities.User) error {
	err := r.db.Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id uuid.UUID, deleteBy uuid.UUID) error {
	err := r.db.Model(&entities.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deleteBy,
	}).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) IsEmailExists(email string) (bool, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) IsPhoneExists(phone string) (bool, error) {
	var user entities.User
	err := r.db.Where("phone_number = ?", phone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).Not("id=?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func (r *userRepository) IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error) {
	var user entities.User
	err := r.db.Where("phone_number = ?", phone).Not("id=?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, nil
	}

	return true, nil
}
