package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func SignAccessToken(user *entity.User, secret string) (*jwt.RegisteredClaims, string, error) {
	secretKey := []byte(secret)

	claims := &jwt.RegisteredClaims{
		Issuer:    "task-management-api",
		Subject:   user.ID.Hex(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign, err := token.SignedString(secretKey)

	return claims, sign, err
}

func ParseAccessToken(tokenString string, secret string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entity.ErrAuthAccessTokenInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, entity.ErrAuthAccessTokenExpired
		}
		return nil, entity.ErrAuthAccessTokenInvalid
	}

	claim, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, entity.ErrAuthAccessTokenInvalid
	}

	return claim, nil
}

func SignRefreshToken(user *entity.User, secret string) (*jwt.RegisteredClaims, string, error) {
	secretKey := []byte(secret)

	claims := &jwt.RegisteredClaims{
		Issuer:    "task-management-api",
		Subject:   user.ID.Hex(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		ID:        bson.NewObjectID().Hex(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign, err := token.SignedString(secretKey)

	return claims, sign, err
}

func ParseRefreshToken(tokenString string, secret string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entity.ErrAuthRefreshTokenInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, entity.ErrAuthRefreshTokenExpired
		}
		return nil, entity.ErrAuthRefreshTokenInvalid
	}

	claim, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, entity.ErrAuthRefreshTokenInvalid
	}

	return claim, nil
}
