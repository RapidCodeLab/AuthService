package server

import (
	"context"
	"net"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/handlers"
	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	auth_grpc "github.com/RapidCodeLab/AuthService/pkg/auth-grpc"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type server struct {
	http         *http.Server
	grpc         *grpc.Server
	configurator interfaces.Configurator
	PublickKey   []byte
	jwtTokener   interfaces.JWTokener
	userService  interfaces.UserService
}

func NewAuthServer(jwtTokener interfaces.JWTokener,
	configurator interfaces.Configurator,
	userService interfaces.UserService) *server {
	return &server{
		jwtTokener:   jwtTokener,
		configurator: configurator,
		userService:  userService,
	}
}

func (s *server) Start(ctx context.Context) (err error) {

	serverErrors := make(chan error, 1)

	//http server start
	r := mux.NewRouter()

	r.HandleFunc(handlers.SigninPath, func(w http.ResponseWriter, r *http.Request) {
		handlers.Signin(w, r, s.jwtTokener, s.userService)
	})

	r.HandleFunc(handlers.SignupPath, func(w http.ResponseWriter, r *http.Request) {
		handlers.Signup(w, r, s.userService)
	})

	r.HandleFunc(handlers.RefreshTokenPath, func(w http.ResponseWriter, r *http.Request) {
		handlers.RefreshToken(w, r, s.jwtTokener)
	})
	r.HandleFunc(handlers.SignoutPath, handlers.Logout)

	s.http = &http.Server{
		Handler: r,
	}

	listener, err := net.Listen(
		s.configurator.GetHTTPServerListenNetwork(),
		s.configurator.GetHTTPServerListenAddr())
	if err != nil {
		return
	}
	go func() {
		serverErrors <- s.http.Serve(listener)
	}()

	//grpc server start
	grpcListener, err := net.Listen(
		s.configurator.GetGRPCServerListenNetwork(),
		s.configurator.GetGRPCServerListenAddr())
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{}
	s.grpc = grpc.NewServer(opts...)

	grpcServer := NewGRPCServer(s.jwtTokener)

	auth_grpc.RegisterAuthGPRCServer(
		s.grpc,
		grpcServer)

	go func() {
		serverErrors <- s.grpc.Serve(grpcListener)
	}()

	go func() {
		<-ctx.Done()
		s.grpc.GracefulStop()
		err := s.http.Shutdown(ctx)
		if err != nil {
			//log err
		}
	}()

	err = <-serverErrors
	return
}
