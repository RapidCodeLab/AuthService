package userservice

import (
	"context"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	user_grpc "github.com/RapidCodeLab/AuthService/pkg/user-grpc"
	"google.golang.org/grpc"
)

type Service struct {
	client user_grpc.UserGPRCClient
}

func New(ctx context.Context,
	configurator interfaces.Configurator,
) (s *Service, err error) {

	conn, err := grpc.DialContext(ctx,
		configurator.GetGRPCUserServiceAddr())
	if err != nil {
		return
	}

	client := user_grpc.NewUserGPRCClient(conn)

	s = &Service{
		client: client,
	}
	return
}

func (s *Service) GetUser(ctx context.Context,
	email, password string) (u interfaces.User, err error) {

	req := &user_grpc.UserRequest{
		Email:    email,
		Password: password,
	}

	res, err := s.client.GetUser(ctx, req)
	if err != nil {
		return
	}

	u.ID = res.GetId()
	u.Email = res.GetEmail()
	u.Status = interfaces.UserStatus(res.GetUserStatus())

	for _, role := range res.GetUserRoles() {
		u.Roles = append(u.Roles, interfaces.UserRole(role))
	}

	return
}
