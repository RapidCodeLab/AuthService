package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RapidCodeLab/AuthService/internal/server"
	"github.com/RapidCodeLab/AuthService/pkg/configurator"
	jwttokener "github.com/RapidCodeLab/AuthService/pkg/jwt-tokener"
	userservice "github.com/RapidCodeLab/AuthService/pkg/services/user"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c := configurator.New()

	jwtTokener, err := jwttokener.New()
	if err != nil {
		log.Fatal(err)
	}

	us, err := userservice.New(ctx, c)
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewAuthServer(jwtTokener, c, us)

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-gracefulStop
		cancel()
	}()

	err = s.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
