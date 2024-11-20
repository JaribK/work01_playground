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
	GetUserById(ctx context.Context, id uuid.UUID) (interface{}, error)
	GetAllUsers(ctx context.Context, page, size int, roleId, isActive string) (models.Pagination, error)
	UpdateUser(user entities.User) error
	DeleteUser(id uuid.UUID, deleteBy uuid.UUID) error
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (s *userUsecase) CreateUser(user entities.User) error {
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

func (s *userUsecase) GetUserById(ctx context.Context, id uuid.UUID) (interface{}, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	var mergedPermissions []models.PermissionDTO
	for _, permission := range user.Role.Permissions {
		mergedPermissions = append(mergedPermissions, models.PermissionDTO{
			ID:           permission.Feature.ID,
			Name:         permission.Feature.Name,
			ParentMenuId: permission.Feature.ParentMenuId,
			MenuIcon:     permission.Feature.MenuIcon,
			MenuNameTh:   permission.Feature.MenuNameTh,
			MenuNameEn:   permission.Feature.MenuNameEn,
			MenuSlug:     permission.Feature.MenuSlug,
			MenuSeqNo:    permission.Feature.MenuSeqNo,
			IsActive:     permission.Feature.IsActive,
			CreateAccess: permission.CreateAccess,
			ReadAccess:   permission.ReadAccess,
			UpdateAccess: permission.UpdateAccess,
			DeleteAccess: permission.DeleteAccess,
		})
	}

	var userDTO []interface{}
	userDTO = append(userDTO, models.ResUserDTO{
		UserID:            user.ID,
		Email:             user.Email,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PhoneNumber:       user.PhoneNumber,
		Avatar:            user.Avatar,
		RoleName:          user.Role.Name,
		RoleLevel:         user.Role.Level,
		TwoFactorAuthUrl:  user.TwoFactorAuthUrl,
		TwoFactorEnabled:  user.TwoFactorEnabled,
		TwoFactorToken:    user.TwoFactorToken,
		TwoFactorVerified: user.TwoFactorVerified,
		Permissions:       mergedPermissions,
	})

	return userDTO, nil
}

func (s *userUsecase) GetAllUsers(ctx context.Context, page, size int, roleId, isActive string) (models.Pagination, error) {
	users, total, err := s.repo.GetAll(ctx, page, size, roleId, isActive)
	if err != nil {
		return models.Pagination{}, err
	}

	var userDTOs []interface{}
	for _, user := range users {
		userDTOs = append(userDTOs, models.ResAllUserDTOs{
			UserID:      user.ID,
			Email:       user.Email,
			FullName:    user.FirstName + " " + user.LastName,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
			Avatar:      user.Avatar,
			RoleName:    user.Role.Name,
		})
	}

	return helpers.Pagiante(page, size, total, userDTOs), nil
}

func (s *userUsecase) UpdateUser(user entities.User) error {
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

func (s *userUsecase) DeleteUser(id uuid.UUID, deleteBy uuid.UUID) error {
	err := s.repo.Delete(id, deleteBy)
	if err != nil {
		return err
	}

	return nil
}
