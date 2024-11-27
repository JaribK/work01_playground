package usecases

import (
	"context"
	"fmt"
	"work01/internal/entities"
	"work01/internal/models"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type RoleUsecase interface {
	CreateRole(role entities.Role) error
	GetRoleById(ctx context.Context, id uuid.UUID) (*entities.Role, error)
	GetAllRolesNoModify() ([]entities.Role, error)
	GetAllRoles(ctx context.Context) ([]models.ResRoleDetails, error)
	GetAllRolesDropdown(ctx context.Context) ([]models.ResRoleDropDown, error)
	UpdateRole(ctx context.Context, role entities.Role) error
	DeleteRole(ctx context.Context, id uuid.UUID, delBy uuid.UUID) error
}

type roleUsecase struct {
	repo repositories.RoleRepository
}

func NewRoleUsecase(repo repositories.RoleRepository) RoleUsecase {
	return &roleUsecase{repo: repo}
}

func (s *roleUsecase) CreateRole(role entities.Role) error {
	if role.Name == "" {
		return fmt.Errorf("role name cannot be empty")
	}
	if err := s.repo.Create(&role); err != nil {
		return err
	}
	return nil
}

func (s *roleUsecase) GetRoleById(ctx context.Context, id uuid.UUID) (*entities.Role, error) {
	role, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *roleUsecase) GetAllRolesNoModify() ([]entities.Role, error) {
	roles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roleUsecase) GetAllRoles(ctx context.Context) ([]models.ResRoleDetails, error) {
	roles, err := s.repo.GetAllModify(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roleUsecase) GetAllRolesDropdown(ctx context.Context) ([]models.ResRoleDropDown, error) {
	roles, err := s.repo.GetAllModify(ctx)
	if err != nil {
		return nil, err
	}

	var roleRes []models.ResRoleDropDown
	for _, role := range roles {
		roleRes = append(roleRes, models.ResRoleDropDown{
			RoleID:   role.RoleID,
			RoleName: role.RoleName,
		})
	}

	return roleRes, nil
}

func (s *roleUsecase) UpdateRole(ctx context.Context, role entities.Role) error {
	err := s.repo.Update(ctx, &role)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleUsecase) DeleteRole(ctx context.Context, id uuid.UUID, delBy uuid.UUID) error {
	role, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if role.Name == "Super Administrator" {
		return fmt.Errorf("can't remove role super administrator")
	}

	err = s.repo.Delete(id, delBy)
	if err != nil {
		return err
	}

	return nil
}
