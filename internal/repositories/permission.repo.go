package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

type PermissionRepository interface {
	Create(permission *entities.Permission) error
	GetById(id uuid.UUID) (*entities.Permission, error)
	GetAll() ([]entities.Permission, error)
	Update(permission *entities.Permission) error
	Delete(id uuid.UUID) error
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) Create(permission *entities.Permission) error {
	err := r.db.Create(&permission).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *permissionRepository) GetById(id uuid.UUID) (*entities.Permission, error) {
	var permission entities.Permission
	err := r.db.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetAll() ([]entities.Permission, error) {
	var permissions []entities.Permission
	err := r.db.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *permissionRepository) Update(permission *entities.Permission) error {
	err := r.db.Where("id=?", permission.ID).Updates(&permission).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *permissionRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&entities.Permission{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
