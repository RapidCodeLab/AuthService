package mockuserservice

import (
	"context"
	"errors"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	"github.com/google/uuid"
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

	u, ok := s.storage[email]
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

	u.ID = uuid.NewString()
	u.Email = email
	u.Status = interfaces.UserStatusNew
	u.Roles = append(u.Roles, interfaces.UserRole(role))

	s.storage[u.Email] = u
	return
}
