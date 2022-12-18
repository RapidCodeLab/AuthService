package server

import (
	"context"
	"net"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/handlers"
	"github.com/RapidCodeLab/AuthService/internal/interfaces"
	"github.com/RapidCodeLab/AuthService/pkg/authgrpcserver"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	LoginPath        = ""
	SignupPath       = ""
	RefreshTokenPath = ""
	LogoutPath       = ""
)

type server struct {
	http       *http.Server
	grpc       *grpc.Server
	PublickKey []byte
	jwtTokener interfaces.JWTokener
}

func NewAuthServer(jwtTokener interfaces.JWTokener) *server {
	return &server{
		jwtTokener: jwtTokener,
	}
}

func (s *server) Start(ctx context.Context) (err error) {

	serverErrors := make(chan error, 1)

	//http server start
	r := mux.NewRouter()

	r.HandleFunc(LoginPath, func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r, s.jwtTokener)
	})
	r.HandleFunc(SignupPath, handlers.Signup)
	r.HandleFunc(RefreshTokenPath, handlers.RefreshToken)
	r.HandleFunc(LogoutPath, handlers.Logout)

	s.http = &http.Server{
		Handler: r,
	}

	listener, err := net.Listen("", "")
	if err != nil {
		return
	}
	go func() {
		serverErrors <- s.http.Serve(listener)
	}()

	//grpc server start
	grpcListener, err := net.Listen("", "")
	if err != nil {
		return
	}
	opts := []grpc.ServerOption{}
	s.grpc = grpc.NewServer(opts...)

	grpcServer := NewGRPCServer(s.jwtTokener)

	authgrpcserver.RegisterAuthServer(
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

func (s *server) Stop() (err error) {
	return
}
