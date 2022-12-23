package server

import (
	"context"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	auth_grpc "github.com/RapidCodeLab/AuthService/pkg/auth-grpc"
)

type grpcServer struct {
	jwtTokener interfaces.JWTokener
	auth_grpc.UnimplementedAuthGPRCServer
}

func NewGRPCServer(jwtTokener interfaces.JWTokener) *grpcServer {
	return &grpcServer{
		jwtTokener: jwtTokener,
	}
}

func (gs *grpcServer) GetPublicKey(
	ctx context.Context,
	empty *auth_grpc.Empty) (
	r *auth_grpc.Response,
	err error) {

	r.PublicKey = gs.jwtTokener.GetPublicKey()
	return
}
