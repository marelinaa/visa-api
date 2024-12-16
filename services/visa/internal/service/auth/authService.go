// Package auth ...
package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/marelinaa/visa-api/services/visa/internal/domain"
	"github.com/marelinaa/visa-api/services/visa/internal/service"

	"github.com/go-playground/validator/v10"
)

// Repository ...
type UserRepo interface {
	Create(ctx context.Context, req domain.User) error
	SignIn(ctx context.Context, userSignInInput domain.SignInInput) (int64, string, error)
}

type Cache interface {
	Set(key string, value interface{}, exp time.Duration) error
}

type AuthService struct {
	userRepo UserRepo
	cache    Cache
	validate *validator.Validate
}

// NewService creates a new Service with the given repository and Redis client.
func NewAuthService(psqlRepo UserRepo, redisRepo Cache) *AuthService {
	authService := &AuthService{
		userRepo: psqlRepo,
		cache:    redisRepo,
		validate: validator.New(),
	}

	return authService
}

// SignUp ...
func (s *AuthService) SignUp(ctx context.Context, userInput domain.SignUpInput) error {
	var profile domain.User

	if err := service.ValidateInput(userInput, s.validate); err != nil {

		return err
	}

	hashedPassword, err := service.HashPassword(userInput.Password)
	if err != nil {

		return domain.ErrHashingPassword
	}

	profile.FirstName = userInput.FirstName
	profile.LastName = userInput.LastName
	profile.PhoneNumber = userInput.PhoneNumber
	profile.Email = userInput.Email
	profile.PasswordHash = hashedPassword
	profile.Role = "applicant" //todo: change giving roles

	return s.userRepo.Create(ctx, profile)
}

// SignIn ...
func (s *AuthService) SignIn(ctx context.Context, userSignInInput domain.SignInInput) (service.Token, error) {
	if err := service.ValidateInput(userSignInInput, s.validate); err != nil {

		return service.Token{}, err
	}

	userID, hash, err := s.userRepo.SignIn(ctx, userSignInInput)
	if err != nil {

		return service.Token{}, err
	}

	if !service.VerifyPassword(userSignInInput.Password, hash) {

		return service.Token{}, domain.ErrUnauthorized
	}

	// Generate tokens
	var token service.Token
	token.Access, token.Refresh, err = service.GenerateToken(userID)
	if err != nil {

		return service.Token{}, err
	}

	accessRefresh := time.Minute * 15
	err = s.cache.Set(
		"access_token:"+fmt.Sprintf("%d", userID), token.Access, accessRefresh)
	if err != nil {

		return service.Token{}, err
	}

	refreshExp := time.Hour * 24
	err = s.cache.Set(
		"refresh_token:"+fmt.Sprintf("%d", userID), token.Refresh, refreshExp)
	if err != nil {

		return service.Token{}, err
	}

	return token, nil
}
