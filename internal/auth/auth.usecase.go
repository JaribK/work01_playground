package auth

import (
	"fmt"
	"work01/internal/entities"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase interface {
	Login(email, password string) (*entities.User, string, error)
}

type AuthUsecase struct {
	repo AuthRepository
}

func NewAuthUsecase(repo AuthRepository) authUsecase {
	return &AuthUsecase{repo: repo}
}

func (s *AuthUsecase) Login(email, password string) (*entities.User, string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("email or password Incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("email or password Incorrect")
	}

	token, err := GenerateToken(*user)
	if err != nil {
		return nil, "", fmt.Errorf("could not generate token: %v", err)
	}

	return user, token, nil

}
