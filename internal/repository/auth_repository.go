package repository

import (
	"context"

	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
)

type authRepository struct {
	mc *bootstrap.MongoCollections
}

func NewAuthHandler(mc *bootstrap.MongoCollections) entity.AuthRepository {
	return &authRepository{mc: mc}
}

func (r *authRepository) Register(ctx context.Context, registerUser *entity.User) (*entity.User, error) {
	if _, err := r.mc.Users.InsertOne(ctx, registerUser); err != nil {
		return nil, err
	}

	return registerUser, nil
}
