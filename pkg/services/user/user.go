package userservice

import (
	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	user_grpc "github.com/RapidCodeLab/AuthService/pkg/user-grpc"
)

type Service struct {
	client *user_grpc.UserGPRCClient
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetUser(user, password string) (u interfaces.User, err error) {

	return
}
