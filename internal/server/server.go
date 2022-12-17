package server

import (
	"net"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/handlers"
	"github.com/gorilla/mux"
)

const (
	LoginPath        = ""
	SignupPath       = ""
	RefreshTokenPath = ""
	LogoutPath       = ""
)

type server struct {
	http *http.Server
}

func New() *server {
	return &server{}
}

func (s *server) Start() (err error) {

	r := mux.NewRouter()

	r.HandleFunc(LoginPath, handlers.Login)
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

	return s.http.Serve(listener)
}

func (s *server) Stop() (err error) {
	return
}
