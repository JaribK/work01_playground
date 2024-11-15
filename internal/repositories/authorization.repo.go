package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authorizationRepository struct {
	db *gorm.DB
}

type AuthorizationRepository interface {
	Create(auth *entities.Authorization) error
	GetById(id uuid.UUID) (*entities.Authorization, error)
	GetAll() ([]entities.Authorization, error)
	Update(auth *entities.Authorization) error
	Delete(id uuid.UUID, deleteBy uuid.UUID) error
	GetUserById(id uuid.UUID) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckAuthorizationByUserID(id uuid.UUID) bool
	GetAuthorizationByUserID(id uuid.UUID) (*entities.Authorization, error)
	DeleteAuthorizationByUserId(id uuid.UUID) error
	GetAuthorizationByRefreshToken(refreshToken string) (*entities.Authorization, error)
}

func NewAuthorizationRepository(db *gorm.DB) AuthorizationRepository {
	return &authorizationRepository{db: db}
}

func (r *authorizationRepository) Create(auth *entities.Authorization) error {
	err := r.db.Create(&auth).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *authorizationRepository) GetById(id uuid.UUID) (*entities.Authorization, error) {
	var auth entities.Authorization
	err := r.db.First(&auth, id).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *authorizationRepository) GetAll() ([]entities.Authorization, error) {
	var auths []entities.Authorization
	err := r.db.Find(&auths).Error
	if err != nil {
		return nil, err
	}

	return auths, nil
}

func (r *authorizationRepository) Update(auth *entities.Authorization) error {
	err := r.db.Where("id=?", auth.ID).Updates(&auth).Error
	if err != nil {
		return err
	}

	return nil
}

// delete with path
func (r *authorizationRepository) Delete(id uuid.UUID, deleteBy uuid.UUID) error {
	err := r.db.Model(&entities.Authorization{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deleteBy,
	}).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&entities.Authorization{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *authorizationRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authorizationRepository) GetUserById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authorizationRepository) CheckAuthorizationByUserID(id uuid.UUID) bool {
	var auth entities.Authorization
	if err := r.db.Where("user_id = ?", id).First(&auth).Error; err != nil {
		return false
	}
	return true
}

func (r *authorizationRepository) GetAuthorizationByUserID(id uuid.UUID) (*entities.Authorization, error) {
	var auth entities.Authorization
	if err := r.db.Where("user_id = ?", id).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (r *authorizationRepository) GetAuthorizationByRefreshToken(refreshToken string) (*entities.Authorization, error) {
	var auth entities.Authorization
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

// for logout
func (r *authorizationRepository) DeleteAuthorizationByUserId(id uuid.UUID) error {
	err := r.db.Model(&entities.Authorization{}).Where("user_id = ?", id).Updates(map[string]interface{}{
		"access_token":  "",
		"refresh_token": "",
	}).Error
	if err != nil {
		return err
	}
	return nil
}
