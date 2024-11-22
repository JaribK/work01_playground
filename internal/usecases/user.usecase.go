package usecases

import (
	"context"
	"errors"
	"fmt"
	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/models"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(user entities.User) error
	GetUserById(ctx context.Context, id uuid.UUID) (*models.ResUserDTO, error)
	GetAllUsers(ctx context.Context, page, size int, roleId, isActive string) (models.Pagination, error)
	UpdateUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (s *userUsecase) CreateUser(user entities.User) error {
	if err := s.CheckVariableToCreate(user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password); err != nil {
		return err
	}

	if err := s.repo.Create(&user); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) GetUserById(ctx context.Context, id uuid.UUID) (*models.ResUserDTO, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userUsecase) GetAllUsers(ctx context.Context, page, size int, roleId, isActive string) (models.Pagination, error) {
	users, total, err := s.repo.GetAll(ctx, page, size, roleId, isActive)
	if err != nil {
		return models.Pagination{}, err
	}

	return helpers.Pagiante(page, size, total, users), nil
}

func (s *userUsecase) UpdateUser(ctx context.Context, user entities.User) error {
	if err := s.CheckVariableToUpdate(user.ID, user.Email, user.PhoneNumber); err != nil {
		return err
	}

	if err := s.repo.Update(ctx, &user); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) DeleteUser(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if user.RoleName == "Super Administrator" {
		return fmt.Errorf("can't remove user that's have role super administrator")
	}

	if err := s.repo.Delete(ctx, id, deleteBy); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) CheckVariableToCreate(firstName string, lastName string, email string, phoneNumber string, password string) error {
	if firstName == "" || lastName == "" || email == "" || phoneNumber == "" || password == "" {
		return fmt.Errorf("please fill all theese field -> firstName, lastName, email, phoneNumber and password")
	}

	emailExists, err := s.repo.IsEmailExists(email)
	if err != nil {
		return err
	}

	if emailExists {
		return fmt.Errorf("email already exists")
	}

	phoneExists, err := s.repo.IsPhoneExists(phoneNumber)
	if err != nil {
		return err
	}

	if phoneExists {
		return fmt.Errorf("phone already exists")
	}

	return nil
}

func (s *userUsecase) CheckVariableToUpdate(userId uuid.UUID, email string, phoneNumber string) error {
	if email != "" {
		Exists, err := s.repo.IsEmailExistsForUpdate(email, userId)
		if err != nil {
			return err
		}

		if Exists {
			return errors.New("email already exists")
		}
	}

	if phoneNumber != "" {
		Exists, err := s.repo.IsPhoneExistsForUpdate(phoneNumber, userId)
		if err != nil {
			return err
		}

		if Exists {
			return errors.New("phone already exists")
		}
	}

	return nil
}
