package auth

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/big"
	"notes/model"
	"notes/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
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

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generatePasswordSalt() string {
	number, err := rand.Int(rand.Reader, big.NewInt(8))
	if err != nil {
		return ""
	}
	length := number.Int64() + 18
	buff := make([]byte, length)
	_, err = rand.Read(buff)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", buff)[:length]
}

func (u *AuthUsecases) CreateUser(user model.User) (int, error) {
	user.Salt = generatePasswordSalt()
	user.Password = generatePasswordHash(user.Password, user.Salt)

	return u.repository.CreateUser(user)
}

func (u *AuthUsecases) GenerateToken(username, password string) (string, error) {
	user, err := u.repository.GetUser(username)
	if err != nil {
		return "", err
	}
	if user.Password != generatePasswordHash(password, user.Salt) {
		return "", errors.New("password or login incoret")
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

func (u *AuthUsecases) ParseTokenToUserId(accessToken string) (int, error) {
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
