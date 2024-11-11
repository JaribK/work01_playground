package usecases

import (
	"errors"
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
	emailExists, err := s.repo.IsEmailExists(user.Email)
	if err != nil {
		return err
	}

	if emailExists {
		return errors.New("email already exists")
	}

	phoneExists, err := s.repo.IsPhoneExists(user.PhoneNumber)
	if err != nil {
		return err
	}

	if phoneExists {
		return errors.New("phone already exists")
	}

	err = s.repo.Create(&user)
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
	if user.Email != "" {
		Exists, err := s.repo.IsEmailExistsForUpdate(user.Email, user.ID)
		if err != nil {
			return err
		}

		if Exists {
			return errors.New("email already exists")
		}
	}

	if user.PhoneNumber != "" {
		Exists, err := s.repo.IsPhoneExistsForUpdate(user.PhoneNumber, user.ID)
		if err != nil {
			return err
		}

		if Exists {
			return errors.New("phone already exists")
		}
	}

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
