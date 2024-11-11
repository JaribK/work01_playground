package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

type RoleRepository interface {
	GetById(id uuid.UUID) (*entities.Role, error)
	GetAll() ([]entities.Role, error)
	Create(role *entities.Role) error
	Update(role *entities.Role) error
	Delete(id uuid.UUID) error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetById(id uuid.UUID) (*entities.Role, error) {
	var roleOjb entities.Role
	err := r.db.First(&roleOjb, id).Error
	if err != nil {
		return nil, err
	}
	return &roleOjb, nil
}

func (r *roleRepository) GetAll() ([]entities.Role, error) {
	var roleOjbs []entities.Role
	err := r.db.Find(&roleOjbs).Error
	if err != nil {
		return nil, err
	}
	return roleOjbs, nil
}

func (r *roleRepository) Create(role *entities.Role) error {
	err := r.db.Create(role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Update(role *entities.Role) error {
	err := r.db.Where("id=?", role.ID).Updates(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&entities.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
