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
	GetAllRoles(ctx context.Context) (interface{}, error)
	GetAllRolesDropdown(ctx context.Context) (interface{}, error)
	UpdateRole(role entities.Role) error
	DeleteRole(id uuid.UUID, delBy uuid.UUID) error
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
	err := s.repo.Create(&role)
	if err != nil {
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

func (s *roleUsecase) GetAllRoles(ctx context.Context) (interface{}, error) {
	roles, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var roleRes []interface{}
	for _, role := range roles {
		roleRes = append(roleRes, models.ResRoleDetails{
			RoleID:     role.ID,
			RoleName:   role.Name,
			RoleLevel:  role.Level,
			NumberUser: int32(len(role.Users)),
		})
	}

	return roleRes, nil
}

func (s *roleUsecase) GetAllRolesDropdown(ctx context.Context) (interface{}, error) {
	roles, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var roleRes []interface{}
	for _, role := range roles {
		roleRes = append(roleRes, models.ResRoleDropDown{
			RoleID:   role.ID,
			RoleName: role.Name,
		})
	}

	return roleRes, nil
}

func (s *roleUsecase) UpdateRole(role entities.Role) error {
	err := s.repo.Update(&role)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleUsecase) DeleteRole(id uuid.UUID, delBy uuid.UUID) error {
	err := s.repo.Delete(id, delBy)
	if err != nil {
		return err
	}

	return nil
}
