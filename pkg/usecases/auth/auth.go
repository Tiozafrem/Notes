package auth

import (
	"crypto/sha1"
	"fmt"
	"notes/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const (
	salt       = "lkjahsdoi12389jhnduoi37asd"
	signingKey = "ljashdkjashdiuawieu7y192873sad"
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
