package mockuserservice

import (
	"context"
	"errors"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
)

type Service struct {
	storage map[string]interfaces.User
}

func New(ctx context.Context,
	configurator interfaces.Configurator,
) (s *Service, err error) {

	s = &Service{
		storage: make(map[string]interfaces.User),
	}
	return
}

func (s *Service) GetUser(ctx context.Context,
	email, password string) (u interfaces.User, err error) {

	u, ok := s.storage[u.Email]
	if !ok {
		err = errors.New("user not found")
	}
	return
}

func (s *Service) CreateUser(
	ctx context.Context,
	email,
	password string,
	role int) (u interfaces.User, err error) {

	s.storage[u.Email] = u
	return
}
