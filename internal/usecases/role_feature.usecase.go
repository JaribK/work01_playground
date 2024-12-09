package usecases

import (
	"context"
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type (
	RoleFeatureUsecase interface {
		CreateRoleFeature(rolePermission entities.RoleFeature) error
		GetRoleFeatureById(ctx context.Context, id uuid.UUID) (*entities.RoleFeature, error)
		GetAllRoleFeatures(ctx context.Context) ([]entities.RoleFeature, error)
		UpdateRoleFeature(rolePermission entities.RoleFeature) error
		DeleteRoleFeature(id uuid.UUID) error
	}

	rolePermissionUsecase struct {
		repo repositories.RoleFeatureRepository
	}
)

func NewRoleFeatureUsecase(repo repositories.RoleFeatureRepository) RoleFeatureUsecase {
	return &rolePermissionUsecase{repo: repo}
}

func (s *rolePermissionUsecase) CreateRoleFeature(rolePermission entities.RoleFeature) error {
	if err := s.repo.Create(&rolePermission); err != nil {
		return err
	}
	return nil
}

func (s *rolePermissionUsecase) GetRoleFeatureById(ctx context.Context, id uuid.UUID) (*entities.RoleFeature, error) {
	rolePermission, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (s *rolePermissionUsecase) GetAllRoleFeatures(ctx context.Context) ([]entities.RoleFeature, error) {
	rolePermissions, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (s *rolePermissionUsecase) UpdateRoleFeature(rolePermission entities.RoleFeature) error {
	if err := s.repo.Update(&rolePermission); err != nil {
		return err
	}

	return nil
}

func (s *rolePermissionUsecase) DeleteRoleFeature(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
