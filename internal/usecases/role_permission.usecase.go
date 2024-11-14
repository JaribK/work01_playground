package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type RolePermissionUsecase interface {
	CreateRolePermission(rolePermission entities.RolePermission) error
	GetRolePermissionById(id uuid.UUID) (*entities.RolePermission, error)
	GetAllRolePermissions() ([]entities.RolePermission, error)
	UpdateRolePermission(rolePermission entities.RolePermission) error
	DeleteRolePermission(id uuid.UUID) error
}

type rolePermissionUsecase struct {
	repo repositories.RolePermissionRepository
}

func NewRolePermissionUsecase(repo repositories.RolePermissionRepository) RolePermissionUsecase {
	return &rolePermissionUsecase{repo: repo}
}

func (s *rolePermissionUsecase) CreateRolePermission(rolePermission entities.RolePermission) error {
	err := s.repo.Create(&rolePermission)
	if err != nil {
		return err
	}
	return nil
}

func (s *rolePermissionUsecase) GetRolePermissionById(id uuid.UUID) (*entities.RolePermission, error) {
	rolePermission, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (s *rolePermissionUsecase) GetAllRolePermissions() ([]entities.RolePermission, error) {
	rolePermissions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (s *rolePermissionUsecase) UpdateRolePermission(rolePermission entities.RolePermission) error {
	err := s.repo.Update(&rolePermission)
	if err != nil {
		return err
	}

	return nil
}

func (s *rolePermissionUsecase) DeleteRolePermission(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
