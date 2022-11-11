package main

import (
	"context"
	"go-clean-architecture/dep"
	pkg_config "go-clean-architecture/pkg/config"
	pkg_sentry "go-clean-architecture/pkg/sentry"
	pkg_validator "go-clean-architecture/pkg/validator"
	"go-clean-architecture/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Config
	pkg_config.NewConfig()

	// Validator
	pkg_validator.NewValidator()

	// Sentry
	pkg_sentry.NewSentry()

	// Server
	server, err := dep.InitializeServer()
	if err != nil {
		utils.CaptureError(err)
	}
	go func() {
		err := server.Run()
		if err != nil {
			utils.CaptureError(err)
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
