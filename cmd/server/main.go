package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/RapidCodeLab/AuthService/internal/server"
)

func main() {

	ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)

  s := server.New()
  

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
