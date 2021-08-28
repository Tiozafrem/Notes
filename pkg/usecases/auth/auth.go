package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"notes/model"
	"notes/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "lkjahsdoi12389jhnduoi37asd"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthUsecases struct {
	repository repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthUsecases(repository repository.Authorization) *AuthUsecases {
	return &AuthUsecases{repository: repository}
}

func generatePasswordHas(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (authUsecases *AuthUsecases) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHas(user.Password)
	return authUsecases.repository.CreateUser(user)
}

func (authUsecases *AuthUsecases) GenerateToken(username, password string) (string, error) {
	user, err := authUsecases.repository.GetUser(username, generatePasswordHas(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (authUsecases *AuthUsecases) ParseTokenToUserId(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(signingKey), nil
		})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of tupe *tokenClaims")
	}

	return claims.UserId, nil
}
