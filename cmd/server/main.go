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
	mockskvtorage "github.com/RapidCodeLab/AuthService/pkg/mocks/kv-storage"
	mockuserservice "github.com/RapidCodeLab/AuthService/pkg/mocks/user-service"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c, err := configurator.New()
	if err != nil {
		log.Fatal(err)
	}

	kv := mockskvtorage.New()

	jwtTokener, err := jwttokener.New(kv)
	if err != nil {
		log.Fatal(err)
	}

	us, err := mockuserservice.New(ctx, c)
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
