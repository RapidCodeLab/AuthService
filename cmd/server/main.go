package main

import (
	"context"
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

	s := server.New(jwtTokener)

	go func() {
		s.Start()
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-gracefulStop
		s.Stop()
		cancel()
	}()

	//start server with context
}
