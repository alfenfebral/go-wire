package main

import (
	"context"
	"go-clean-architecture/dep"
	pkg_config "go-clean-architecture/pkg/config"
	pkg_validator "go-clean-architecture/pkg/validator"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	// Config
	pkg_config.NewConfig()

	// Validator
	pkg_validator.NewValidator()

	// Server
	server, err := dep.InitializeServer()
	if err != nil {
		logrus.Error(err)
	}
	go func() {
		err := server.Run()
		if err != nil {
			logrus.Error(err)
		}
	}()

	// catch shutdown
	done := make(chan bool, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		// graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.GracefulStop(ctx, done)
	}()

	// wait for graceful shutdown
	<-done
}
