package usecases

import (
	"errors"
	"fmt"
	"work01/internal/entities"
	"work01/internal/repositories"
	"work01/pkg"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type UserUseCase interface {
	CreateUser(user entities.User) error
	GetUserById(id uuid.UUID) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(id uuid.UUID) error
	Login(email, password string) (*entities.User, string, error)
}

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entities.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.PhoneNumber == "" || user.Password == "" {
		return fmt.Errorf("please fill all theese field -> firstName, lastName, email, phoneNumber and password")
	}

	emailExists, err := s.repo.IsEmailExists(user.Email)
	if err != nil {
		return err
	}

	if emailExists {
		return fmt.Errorf("email already exists")
	}

	phoneExists, err := s.repo.IsPhoneExists(user.PhoneNumber)
	if err != nil {
		return err
	}

	if phoneExists {
		return fmt.Errorf("phone already exists")
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

func (s UserService) Login(email, password string) (*entities.User, string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("email or password Incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("email or password Incorrect")
	}

	token, err := pkg.GenerateToken(*user)
	if err != nil {
		return nil, "", fmt.Errorf("could not generate token: %v", err)
	}

	return user, token, nil

}