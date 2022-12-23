package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RapidCodeLab/AuthService/internal/server"
	jwttokener "github.com/RapidCodeLab/AuthService/pkg/jwt-tokener"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	jwtTokener := jwttokener.New()

	s := server.NewAuthServer(jwtTokener, nil)

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-gracefulStop
		cancel()
	}()

	err := s.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
