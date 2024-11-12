package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type RolePermissionUseCase interface {
	CreateRolePermission(rolePermission entities.RolePermission) error
	GetRolePermissionById(id uuid.UUID) (*entities.RolePermission, error)
	GetAllRolePermissions() ([]entities.RolePermission, error)
	UpdateRolePermission(rolePermission entities.RolePermission) error
	DeleteRolePermission(id uuid.UUID) error
}

type RolePermissionService struct {
	repo repositories.RolePermissionRepository
}

func NewRolePermissionService(repo repositories.RolePermissionRepository) RolePermissionUseCase {
	return &RolePermissionService{repo: repo}
}

func (s *RolePermissionService) CreateRolePermission(rolePermission entities.RolePermission) error {
	err := s.repo.Create(&rolePermission)
	if err != nil {
		return err
	}
	return nil
}

func (s *RolePermissionService) GetRolePermissionById(id uuid.UUID) (*entities.RolePermission, error) {
	rolePermission, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (s *RolePermissionService) GetAllRolePermissions() ([]entities.RolePermission, error) {
	rolePermissions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (s *RolePermissionService) UpdateRolePermission(rolePermission entities.RolePermission) error {
	err := s.repo.Update(&rolePermission)
	if err != nil {
		return err
	}

	return nil
}

func (s *RolePermissionService) DeleteRolePermission(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
