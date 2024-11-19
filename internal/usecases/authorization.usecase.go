package usecases

import (
	"errors"
	"fmt"
	"time"
	"work01/internal/auth"
	"work01/internal/entities"
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
	err := s.repo.Create(&auth)
	if err != nil {
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

func (s *authorizationUsecase) GetAllAuthorizations() ([]entities.Authorization, error) {
	auths, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return auths, nil
}

func (s *authorizationUsecase) UpdateAuthorization(auth entities.Authorization) error {
	err := s.repo.Update(&auth)
	if err != nil {
		return err
	}

	return nil
}

func (s *authorizationUsecase) DeleteAuthorization(id uuid.UUID, delBy uuid.UUID) error {
	err := s.repo.Delete(id, delBy)
	if err != nil {
		return err
	}

	return nil
}

func (s *authorizationUsecase) Login(email, password string) (*entities.User, *models.AuthToken, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, nil, fmt.Errorf("email or password Incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, fmt.Errorf("email or password Incorrect")
	}

	token, err := auth.GenerateToken(user)
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
	token, err := auth.ValidateToken(tokenString)
	if err != nil {
		return fmt.Errorf("token validation failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["exp"] == nil {
		return errors.New("invalid token claims")
	}

	expiration := int64(claims["exp"].(float64))

	// Calculate TTL for Redis
	currentTime := time.Now().Unix()
	ttl := time.Duration(expiration-currentTime) * time.Second
	if ttl <= 0 {
		ttl = 0 // Immediate expiration
	}

	err = s.repo.DeleteAuthorizationByUserId(id, tokenString, ttl)
	if err != nil {
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

	_, err = auth.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	newAccessToken, err := auth.GenerateToken(user)
	authorization.AccessToken = newAccessToken.AccessToken
	if err != nil {
		return "", err
	}

	if err := s.repo.Update(authorization); err != nil {
		return "", err
	}

	return newAccessToken.AccessToken, nil
}
