package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"work01/config"
	"work01/internal/entities"

	"github.com/golang-jwt/jwt/v5"
)

func loadPrivateKey() *rsa.PrivateKey {
	if err := config.LoadConfig(); err != nil {
		return nil
	}

	cfg := config.ReadInConfig()

	keyData, err := ioutil.ReadFile(cfg.PRIVATE_KEY)
	if err != nil {
		log.Fatalf("error reading private key: %v", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	return privateKey
}

func loadPublicKey() *rsa.PublicKey {
	if err := config.LoadConfig(); err != nil {
		return nil
	}

	cfg := config.ReadInConfig()

	keyData, err := ioutil.ReadFile(cfg.PUBLIC_KEY)
	if err != nil {
		log.Fatalf("error reading private key: %v", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	return publicKey
}

func GenerateToken(user entities.User) (string, error) {
	privateKey := loadPrivateKey()

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["userId"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	publicKey := loadPublicKey()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not parse token: %v", err)
	}

	return token, err
}
