package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authRepository entity.AuthRepository
	timeout        time.Duration
	cost           int
}

func NewAuthUsecase(authRepository entity.AuthRepository, timeout time.Duration, config *bootstrap.Config) entity.AuthUsecase {
	cost, _ := strconv.Atoi(config.HASH_COST)
	return &authUsecase{authRepository: authRepository, timeout: timeout, cost: cost}
}

func (u *authUsecase) Register(req *entity.RegisterRequest) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), u.cost)
	if err != nil {
		return nil, err
	}

	registerUser := entity.User{
		ID:             bson.NewObjectID(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		HashedPassword: string(hashPassword),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return u.authRepository.Register(ctx, &registerUser)
}
