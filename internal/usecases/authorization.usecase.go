package usecases

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthorizationUsecase interface {
		CreateAuthorization(auth entities.Authorization) error
		GetAuthorizationById(id uuid.UUID) (*entities.Authorization, error)
		GetAllAuthorizations() ([]entities.Authorization, error)
		GetUserDataById(id uuid.UUID) (*entities.ResUserDTO, error)
		UpdateAuthorization(auth entities.Authorization) error
		DeleteAuthorization(id uuid.UUID, delBy uuid.UUID) error
		Login(email, password string) (*entities.User, *entities.AuthToken, error)
		Logout(id uuid.UUID, token string) error
		RefreshToken(refreshToken string) (string, error)
	}

	authorizationUsecase struct {
		repo repositories.AuthorizationRepository
	}
)

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

func (s *authorizationUsecase) GetUserDataById(id uuid.UUID) (*entities.ResUserDTO, error) {
	user, err := s.repo.GetUserByIdModify(id)
	if err != nil {
		return nil, err
	}

	return user, nil
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

func (s *authorizationUsecase) Login(identifier, password string) (*entities.User, *entities.AuthToken, error) {
	var user *entities.User
	var err error

	if strings.Contains(identifier, "@") {
		user, err = s.repo.GetUserByEmail(identifier)
		if err != nil {
			return nil, nil, fmt.Errorf("user not found")
		}

		if !*user.IsActive {
			return nil, nil, fmt.Errorf("your account was deactivated")
		}
	} else {
		user, err = s.repo.GetUserByPhoneNumber(identifier)
		if err != nil {
			return nil, nil, fmt.Errorf("user not found")
		}
		if !*user.IsActive {
			return nil, nil, fmt.Errorf("your account was deactivated")
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil, fmt.Errorf("email/phoneNumner or password is invalid")
	}

	token, err := helpers.GenerateToken(user)
	if err != nil {
		return nil, nil, fmt.Errorf("could not generate token: %v", err)
	}

	haveauthcheck := s.repo.CheckAuthorizationByUserID(user.ID)

	if haveauthcheck {
		authSelect, err := s.repo.GetAuthorizationByUserID(user.ID)
		if err != nil {
			return nil, nil, err
		}
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
