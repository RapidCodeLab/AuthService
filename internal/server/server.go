package server

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"net"
	"net/http"

	"github.com/RapidCodeLab/AuthService/internal/handlers"
	"github.com/cristalhq/jwt"
	"github.com/gorilla/mux"
)

const (
	LoginPath        = ""
	SignupPath       = ""
	RefreshTokenPath = ""
	LogoutPath       = ""
)

type server struct {
	http       *http.Server
	PublickKey []byte
	jwtBuilder *jwt.TokenBuilder
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

func (s *server) TokenBuilderUpdate() (err error) {

	privateKey, err :=
		ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}

	signer, err := jwt.NewES256(&privateKey.PublicKey, privateKey)
	if err != nil {
		return
	}
	s.jwtBuilder = jwt.NewTokenBuilder(signer)
	return
}
