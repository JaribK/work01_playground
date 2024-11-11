package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type RoleUseCase interface {
	CreateRole(role entities.Role) error
	GetRoleById(id uuid.UUID) (*entities.Role, error)
	GetAllRoles() ([]entities.Role, error)
	UpdateRole(role entities.Role) error
	DeleteRole(id uuid.UUID) error
}

type RoleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleUseCase {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(role entities.Role) error {
	err := s.repo.Create(&role)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) GetRoleById(id uuid.UUID) (*entities.Role, error) {
	role, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *RoleService) GetAllRoles() ([]entities.Role, error) {
	roles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RoleService) UpdateRole(role entities.Role) error {
	err := s.repo.Update(&role)
	if err != nil {
		return err
	}

	return nil
}

func (s *RoleService) DeleteRole(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
