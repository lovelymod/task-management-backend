package utils

import (
	"errors"
	"net/http"

	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetStatusError(err error) int {
	if errors.Is(err, entity.ErrGlobalServerError) {
		return http.StatusInternalServerError
	}
	if errors.Is(err, entity.ErrGlobalNotFound) || errors.Is(err, mongo.ErrNoDocuments) {
		return http.StatusNotFound
	}
	if errors.Is(err, entity.ErrGlobalNotHavePermission) {
		return http.StatusForbidden
	}
	if errors.Is(err, entity.ErrAuthRefreshTokenExpired) ||
		errors.Is(err, entity.ErrAuthRefreshTokenInvalid) ||
		errors.Is(err, entity.ErrAuthRefreshTokenNotProvided) ||
		errors.Is(err, entity.ErrAuthAccessTokenExpired) ||
		errors.Is(err, entity.ErrAuthAccessTokenInvalid) ||
		errors.Is(err, entity.ErrAuthAccessTokenNotProvided) ||
		errors.Is(err, entity.ErrAuthWrongEmailOrPassword) {
		return http.StatusUnauthorized
	}

	return http.StatusBadRequest
}
