package service

import (
	"github.com/marelinaa/visa-api/services/gateway/internal/domain"
)

type GatewayService struct {
	users map[string]string
}

func NewGatewayService(users map[string]string) *GatewayService {
	gatewayService := &GatewayService{
		users: users,
	}

	return gatewayService
}

func (s *GatewayService) SignIn(userSignIn domain.User) error {
	if userSignIn.Email == "" || userSignIn.Password == "" {
		return domain.ErrEmptyInput
	}

	password, ok := s.users[userSignIn.Email]
	if !ok || password != userSignIn.Password {
		return domain.ErrInvalidCredentials
	}

	return nil
}
