package usecases

import (
	"errors"
	"fmt"
	"time"
	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/models"
	"work01/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authorizationUsecase struct {
	repo repositories.AuthorizationRepository
}

type AuthorizationUsecase interface {
	CreateAuthorization(auth entities.Authorization) error
	GetAuthorizationById(id uuid.UUID) (*entities.Authorization, error)
	GetAllAuthorizations() ([]entities.Authorization, error)
	GetUserDataById(id uuid.UUID) (interface{}, error)
	UpdateAuthorization(auth entities.Authorization) error
	DeleteAuthorization(id uuid.UUID, delBy uuid.UUID) error
	Login(email, password string) (*entities.User, *models.AuthToken, error)
	Logout(id uuid.UUID, token string) error
	RefreshToken(refreshToken string) (string, error)
}

func NewAuthorizationUsecase(repo repositories.AuthorizationRepository) AuthorizationUsecase {
	return &authorizationUsecase{repo: repo}
}

func (s *authorizationUsecase) CreateAuthorization(auth entities.Authorization) error {
	if err := s.repo.Create(&auth); err != nil {
		return err
	}
	return nil
}

func (s *authorizationUsecase) GetAuthorizationById(id uuid.UUID) (*entities.Authorization, error) {
	auth, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (s *authorizationUsecase) GetUserDataById(id uuid.UUID) (interface{}, error) {
	user, err := s.repo.GetUserById(id)
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

	userDTO := models.ResUserDTO{
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
	}

	return &userDTO, nil
}

func (s *authorizationUsecase) GetAllAuthorizations() ([]entities.Authorization, error) {
	auths, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return auths, nil
}

func (s *authorizationUsecase) UpdateAuthorization(auth entities.Authorization) error {
	if err := s.repo.Update(&auth); err != nil {
		return err
	}

	return nil
}

func (s *authorizationUsecase) DeleteAuthorization(id uuid.UUID, delBy uuid.UUID) error {
	if err := s.repo.Delete(id, delBy); err != nil {
		return err
	}

	return nil
}

func (s *authorizationUsecase) Login(email, password string) (*entities.User, *models.AuthToken, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, nil, fmt.Errorf("email or password Incorrect")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil, fmt.Errorf("email or password Incorrect")
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return nil, nil, fmt.Errorf("could not generate token: %v", err)
	}

	haveauthcheck := s.repo.CheckAuthorizationByUserID(user.ID)
	authSelect, err := s.repo.GetAuthorizationByUserID(user.ID)
	if err != nil {
		return nil, nil, err
	}

	if haveauthcheck {
		auth := &entities.Authorization{
			ID:           authSelect.ID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			UpdatedBy:    user.ID,
		}

		if err := s.repo.Update(auth); err != nil {
			return nil, nil, err
		}
	} else {
		auth := &entities.Authorization{
			ID:           uuid.New(),
			UserId:       user.ID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			CreatedBy:    user.ID,
		}

		if err := s.repo.Create(auth); err != nil {
			return nil, nil, err
		}
	}

	return user, token, nil

}

func (s *authorizationUsecase) Logout(id uuid.UUID, tokenString string) error {
	token, err := helpers.ValidateToken(tokenString)
	if err != nil {
		return fmt.Errorf("token validation failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["exp"] == nil {
		return errors.New("invalid token claims")
	}

	expiration := int64(claims["exp"].(float64))

	currentTime := time.Now().Unix()
	ttl := time.Duration(expiration-currentTime) * time.Second
	if ttl <= 0 {
		ttl = 0
	}

	if err = s.repo.DeleteAuthorizationByUserId(id, tokenString, ttl); err != nil {
		return err
	}

	return nil
}

func (s *authorizationUsecase) RefreshToken(refreshToken string) (string, error) {
	authorization, err := s.repo.GetAuthorizationByRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetUserById(authorization.UserId)
	if err != nil {
		return "", err
	}

	_, err = helpers.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	newAccessToken, err := helpers.GenerateToken(user)
	authorization.AccessToken = newAccessToken.AccessToken
	if err != nil {
		return "", err
	}

	if err := s.repo.Update(authorization); err != nil {
		return "", err
	}

	return newAccessToken.AccessToken, nil
}
