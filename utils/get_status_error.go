package utils

import (
	"errors"
	"net/http"

	"github.com/lovelymod/task-management-backend/internal/entity"
)

func GetStatusError(err error) int {
	if errors.Is(err, entity.ErrGlobalServerError) {
		return http.StatusInternalServerError
	}
	if errors.Is(err, entity.ErrGlobalNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, entity.ErrAuthRefreshTokenExpired) ||
		errors.Is(err, entity.ErrAuthRefreshTokenInvalid) ||
		errors.Is(err, entity.ErrAuthRefreshTokenNotProvided) ||
		errors.Is(err, entity.ErrAuthAccessTokenExpired) ||
		errors.Is(err, entity.ErrAuthAccessTokenInvalid) ||
		errors.Is(err, entity.ErrAuthAccessTokenNotProvided) ||
		errors.Is(err, entity.ErrAuthThisEmailIsAlreadyUsed) {
		return http.StatusUnauthorized
	}

	return http.StatusBadRequest
}
