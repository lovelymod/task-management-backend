package repository

import (
	"context"

	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type authRepository struct {
	mc *bootstrap.MongoCollections
}

func NewAuthHandler(mc *bootstrap.MongoCollections) entity.AuthRepository {
	return &authRepository{mc: mc}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) ([]entity.User, error) {
	var exitingUser []entity.User

	filter := bson.D{{Key: "email", Value: email}}
	opts := options.Find().SetLimit(1)

	cursor, err := r.mc.Users.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &exitingUser); err != nil {
		return nil, err
	}

	return exitingUser, nil
}

func (r *authRepository) Register(ctx context.Context, registerUser *entity.User) (*entity.User, error) {
	if _, err := r.mc.Users.InsertOne(ctx, registerUser); err != nil {
		return nil, err
	}

	return registerUser, nil
}
