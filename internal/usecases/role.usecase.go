package usecases

import (
	"fmt"
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type RoleUsecase interface {
	CreateRole(role entities.Role) error
	GetRoleById(id uuid.UUID) (*entities.Role, error)
	GetAllRoles() ([]entities.Role, error)
	UpdateRole(role entities.Role) error
	DeleteRole(id uuid.UUID) error
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

func (s *roleUsecase) GetRoleById(id uuid.UUID) (*entities.Role, error) {
	role, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *roleUsecase) GetAllRoles() ([]entities.Role, error) {
	roles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *roleUsecase) UpdateRole(role entities.Role) error {
	err := s.repo.Update(&role)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleUsecase) DeleteRole(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
