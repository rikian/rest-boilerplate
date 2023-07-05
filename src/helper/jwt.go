package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJwtImpl(secretJwt []byte) JWT {
	return &JwtImpl{
		secretJwt: secretJwt,
	}
}

type JWT interface {
	GenerateJWT(data interface{}, expired int) (string, error)
	ClaimJWT(tokenStr string) (jwt.MapClaims, error)
}

type JwtImpl struct {
	secretJwt []byte
}

func (j *JwtImpl) validateJwtToken(t string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	tokenParts := strings.Split(t, ".")
	if len(tokenParts) != 3 {
		err = errors.New("invalid token parts")
		return "", err
	}
	return t, nil
}

func (j *JwtImpl) GenerateJWT(data interface{}, expired int) (string, error) {
	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	})

	return j.validateJwtToken(encrypt.SignedString(j.secretJwt))
}

func (j *JwtImpl) ClaimJWT(tokenStr string) (jwt.MapClaims, error) {
	var err error
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return j.secretJwt, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		err = errors.New("failed claim token")
		return nil, err
	}
}
