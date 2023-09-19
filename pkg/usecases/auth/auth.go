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
	salt            = "lkjahsdoi12389jhnduoi37asd"
	signingKey      = "qrkjk#4#%35FSFJlja#4353KSFjH"
	accessTokenTTL  = 12 * time.Hour
	refreshTokenTTL = 720 * time.Hour
)

type AuthUsecases struct {
	repository repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	DeviceId int `json:"device_id"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
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

func (u *AuthUsecases) NewAccessToken(deviceId int) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		deviceId,
	})

	return accessToken.SignedString([]byte(signingKey))
}

func (u *AuthUsecases) GenerateToken(username, password, nameDevice string) (Tokens, error) {
	var tokens Tokens

	user, err := u.repository.GetUser(username)
	if err != nil {
		return Tokens{}, err
	}

	if user.Password != generatePasswordHash(password, user.Salt) {
		return Tokens{}, errors.New("password or login incoret")
	}

	tokens.RefreshToken, err = u.NewRefreshToken()
	if err != nil {
		return Tokens{}, err
	}

	deviceId, err := u.repository.CreateDevice(model.DeviceUser{
		Name:         nameDevice,
		Description:  nameDevice,
		UserId:       user.Id,
		Expire:       time.Now().Add(refreshTokenTTL),
		RefreshToken: tokens.RefreshToken,
	})
	if err != nil {
		return Tokens{}, err
	}

	tokens.AccessToken, err = u.NewAccessToken(deviceId)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, nil
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
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	user, err := (u.repository.GetUserByDeviceId(claims.DeviceId))
	return user.Id, err
}

func (u *AuthUsecases) NewRefreshToken() (string, error) {
	buf := make([]byte, 46)

	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf)[:46], nil
}

func (u *AuthUsecases) RefreshToken(refreshToken string) (Tokens, error) {
	deviceUser, err := u.repository.GetDeviceByRefreshToken(refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	if deviceUser.Expire.Before(time.Now()) {
		if err := u.repository.DeleteDeviceByDeviceId(deviceUser.Id); err != nil {
			return Tokens{}, err
		}
		return Tokens{}, errors.New("device expire")
	}

	deviceUser.RefreshToken, err = u.NewRefreshToken()
	if err != nil {
		return Tokens{}, err
	}

	deviceUser.Expire = time.Now().Add(refreshTokenTTL)

	if err := u.repository.UpdateRefreshTokenByDevice(deviceUser); err != nil {
		return Tokens{}, err
	}

	var tokens Tokens
	tokens.RefreshToken = deviceUser.RefreshToken
	tokens.AccessToken, err = u.NewAccessToken(deviceUser.Id)
	return tokens, err
}
