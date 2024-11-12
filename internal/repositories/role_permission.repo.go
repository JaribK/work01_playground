package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rolePermissionRepository struct {
	db *gorm.DB
}

type RolePermissionRepository interface {
	GetById(id uuid.UUID) (*entities.RolePermission, error)
	GetAll() ([]entities.RolePermission, error)
	Create(rolePermission *entities.RolePermission) error
	Update(rolePermission *entities.RolePermission) error
	Delete(id uuid.UUID) error
}

func NewRolePermissionRepository(db *gorm.DB) RolePermissionRepository {
	return &rolePermissionRepository{db: db}
}

func (r *rolePermissionRepository) GetById(id uuid.UUID) (*entities.RolePermission, error) {
	var rolePermission entities.RolePermission
	err := r.db.First(&rolePermission, id).Error
	if err != nil {
		return nil, err
	}
	return &rolePermission, nil
}

func (r *rolePermissionRepository) GetAll() ([]entities.RolePermission, error) {
	var rolePermission []entities.RolePermission
	err := r.db.Find(&rolePermission).Error
	if err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (r *rolePermissionRepository) Create(rolePermission *entities.RolePermission) error {
	err := r.db.Create(&rolePermission).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionRepository) Update(rolePermission *entities.RolePermission) error {
	err := r.db.Where("id=?", rolePermission.ID).Updates(&rolePermission).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&entities.RolePermission{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
