package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type PermissionUsecase interface {
	CreatePermission(permission entities.Permission) error
	GetPermissionById(id uuid.UUID) (*entities.Permission, error)
	GetAllPermissions() ([]entities.Permission, error)
	UpdatePermission(permission entities.Permission) error
	DeletePermission(id uuid.UUID) error
}

type permissionUsecase struct {
	repo repositories.PermissionRepository
}

func NewPermissionUsecase(repo repositories.PermissionRepository) PermissionUsecase {
	return &permissionUsecase{repo: repo}
}

func (s *permissionUsecase) CreatePermission(permission entities.Permission) error {
	err := s.repo.Create(&permission)
	if err != nil {
		return err
	}
	return nil
}

func (s *permissionUsecase) GetPermissionById(id uuid.UUID) (*entities.Permission, error) {
	permission, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *permissionUsecase) GetAllPermissions() ([]entities.Permission, error) {
	permissions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (s *permissionUsecase) UpdatePermission(permission entities.Permission) error {
	err := s.repo.Update(&permission)
	if err != nil {
		return err
	}

	return nil
}

func (s *permissionUsecase) DeletePermission(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
