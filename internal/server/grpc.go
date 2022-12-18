package server

import (
	"context"

	"github.com/RapidCodeLab/AuthService/pkg/authgrpcserver"
)

type grpcServer struct {
	authgrpcserver.UnimplementedAuthServer
}

func (gs *grpcServer) GetPublicKey(
	ctx context.Context,
	empty *authgrpcserver.Empty) (
	r *authgrpcserver.Response,
	err error) {
	return
}
