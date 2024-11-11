package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type UserUseCase interface {
	CreateUser(user entities.User) error
	GetUserById(id uuid.UUID) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(id uuid.UUID) error
}

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entities.User) error {
	err := s.repo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserById(id uuid.UUID) (*entities.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(user entities.User) error {
	err := s.repo.Update(&user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
