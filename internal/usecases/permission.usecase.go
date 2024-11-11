package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type PermissionUseCase interface {
	CreatePermission(permission entities.Permission) error
	GetPermissionById(id uuid.UUID) (*entities.Permission, error)
	GetAllPermissions() ([]entities.Permission, error)
	UpdatePermission(permission entities.Permission) error
	DeletePermission(id uuid.UUID) error
}

type PermissionService struct {
	repo repositories.PermissionRepository
}

func NewPermissionService(repo repositories.PermissionRepository) PermissionUseCase {
	return &PermissionService{repo: repo}
}

func (s *PermissionService) CreatePermission(permission entities.Permission) error {
	err := s.repo.Create(&permission)
	if err != nil {
		return err
	}
	return nil
}

func (s *PermissionService) GetPermissionById(id uuid.UUID) (*entities.Permission, error) {
	permission, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *PermissionService) GetAllPermissions() ([]entities.Permission, error) {
	permissions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (s *PermissionService) UpdatePermission(permission entities.Permission) error {
	err := s.repo.Update(&permission)
	if err != nil {
		return err
	}

	return nil
}

func (s *PermissionService) DeletePermission(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
