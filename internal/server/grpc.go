package server

import (
	"context"

	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	"github.com/RapidCodeLab/AuthService/pkg/authgrpcserver"
)

type grpcServer struct {
	jwtTokener interfaces.JWTokener
	authgrpcserver.UnimplementedAuthServer
}

func NewGRPCServer(jwtTokener interfaces.JWTokener) *grpcServer {
	return &grpcServer{
		jwtTokener: jwtTokener,
	}
}

func (gs *grpcServer) GetPublicKey(
	ctx context.Context,
	empty *authgrpcserver.Empty) (
	r *authgrpcserver.Response,
	err error) {

	r.PublicKey = gs.jwtTokener.GetPublicKey()
	return
}
