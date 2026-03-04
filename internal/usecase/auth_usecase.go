package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"github.com/lovelymod/task-management-backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	repo               entity.AuthRepository
	timeout            time.Duration
	cost               int
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewAuthUsecase(authRepository entity.AuthRepository, timeout time.Duration, config *bootstrap.Config) entity.AuthUsecase {
	cost, _ := strconv.Atoi(config.HASH_COST)
	return &authUsecase{
		repo:               authRepository,
		timeout:            timeout,
		cost:               cost,
		accessTokenSecret:  config.ACCESS_TOKEN_SECRET,
		refreshTokenSecret: config.REFRESH_TOKEN_SECRET,
	}
}

func (u *authUsecase) Register(req *entity.RegisterRequest) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	existingUser, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if len(existingUser) > 0 {
		return nil, entity.ErrAuthThisEmailIsAlreadyUsed
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), u.cost)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}

	registerUser := entity.User{
		ID:             bson.NewObjectID(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DisplayName:    req.DisplayName,
		Email:          req.Email,
		HashedPassword: string(hashPassword),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return u.repo.CreateUser(ctx, &registerUser)
}

func (u *authUsecase) Login(req *entity.LoginRequest, clientIP string, userAgent string) (*entity.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	existingUser, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if len(existingUser) == 0 {
		return nil, entity.ErrAuthWrongEmailOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser[0].HashedPassword), []byte(req.Password)); err != nil {
		return nil, entity.ErrAuthWrongEmailOrPassword
	}

	_, signedAccessToken, err := utils.SignAccessToken(&existingUser[0], u.accessTokenSecret)
	if err != nil {
		return nil, err
	}

	claimsRefreshToken, signedRefreshToken, err := utils.SignRefreshToken(&existingUser[0], u.refreshTokenSecret)
	if err != nil {
		return nil, err
	}

	tokenID, err := bson.ObjectIDFromHex(claimsRefreshToken.ID)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}
	savedRefreshToken := entity.RefreshToken{
		ID:        bson.NewObjectID(),
		Token:     signedRefreshToken,
		TokenID:   tokenID,
		UserID:    existingUser[0].ID,
		ClientIP:  clientIP,
		UserAgent: userAgent,
		IsRevoked: false,
		ExpiresAt: claimsRefreshToken.ExpiresAt.Time,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.CreateRefreshToken(ctx, &savedRefreshToken); err != nil {
		return nil, err
	}

	response := entity.LoginResponse{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return &response, nil
}

func (u *authUsecase) RefreshToken(token string, clientIP string, userAgent string) (*entity.RefreshTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	// Check refresh token is valid or not
	_, err := utils.ParseRefreshToken(token, u.refreshTokenSecret)
	if err != nil {
		return nil, err
	}

	// Get refreshToken in db
	existingRefreshToken, err := u.repo.GetRefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if existingRefreshToken.IsRevoked || time.Now().After(existingRefreshToken.ExpiresAt) {
		return nil, entity.ErrAuthRefreshTokenExpired
	}

	// Sign new accessToken
	_, newAccessToken, err := utils.SignAccessToken(existingRefreshToken.User, u.accessTokenSecret)
	if err != nil {
		return nil, err
	}

	// Sign new refreshToken
	newRefreshTokenClaims, newRefreshToken, err := utils.SignRefreshToken(existingRefreshToken.User, u.refreshTokenSecret)
	if err != nil {
		return nil, err
	}

	tokenID, err := bson.ObjectIDFromHex(newRefreshTokenClaims.ID)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}
	userId, err := bson.ObjectIDFromHex(newRefreshTokenClaims.Subject)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}

	savedRefreshToken := entity.RefreshToken{
		ID:        bson.NewObjectID(),
		Token:     newRefreshToken,
		TokenID:   tokenID,
		UserID:    userId,
		ClientIP:  clientIP,
		UserAgent: userAgent,
		IsRevoked: false,
		ExpiresAt: newRefreshTokenClaims.ExpiresAt.Time,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.CreateRefreshToken(ctx, &savedRefreshToken); err != nil {
		return nil, err
	}

	if err := u.repo.RevokeRefreshToken(ctx, token); err != nil {
		return nil, err
	}

	response := entity.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	return &response, nil
}

func (u *authUsecase) Logout(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	return u.repo.RevokeRefreshToken(ctx, token)
}
