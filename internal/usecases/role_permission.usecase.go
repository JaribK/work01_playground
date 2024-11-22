package usecases

import (
	"context"
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type RolePermissionUsecase interface {
	CreateRolePermission(rolePermission entities.RolePermission) error
	GetRolePermissionById(ctx context.Context, id uuid.UUID) (*entities.RolePermission, error)
	GetAllRolePermissions(ctx context.Context) ([]entities.RolePermission, error)
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
	if err := s.repo.Create(&rolePermission); err != nil {
		return err
	}
	return nil
}

func (s *rolePermissionUsecase) GetRolePermissionById(ctx context.Context, id uuid.UUID) (*entities.RolePermission, error) {
	rolePermission, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (s *rolePermissionUsecase) GetAllRolePermissions(ctx context.Context) ([]entities.RolePermission, error) {
	rolePermissions, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (s *rolePermissionUsecase) UpdateRolePermission(rolePermission entities.RolePermission) error {
	if err := s.repo.Update(&rolePermission); err != nil {
		return err
	}

	return nil
}

func (s *rolePermissionUsecase) DeleteRolePermission(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
