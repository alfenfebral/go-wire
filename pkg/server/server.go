package server

import (
	"context"
	pkg_http_server "go-clean-architecture/pkg/server/http"
	"go-clean-architecture/utils"

	"github.com/sirupsen/logrus"
)

type Server interface {
	Run() error
	GracefulStop(ctx context.Context, done chan bool)
}

type ServerImpl struct {
	httpServer pkg_http_server.HTTPServer
}

func NewServer(httpServer pkg_http_server.HTTPServer) (Server, error) {
	return &ServerImpl{
		httpServer: httpServer,
	}, nil
}

// Run server
func (s *ServerImpl) Run() error {
	go func() {
		err := s.httpServer.Run()
		if err != nil {
			utils.CaptureError(err)
		}
	}()

	return nil
}

// GracefulStop server
func (s *ServerImpl) GracefulStop(ctx context.Context, done chan bool) {
	err := s.httpServer.GracefulStop(ctx)
	if err != nil {
		utils.CaptureError(err)
	}

	logrus.Info("gracefully shutdowned")
	done <- true
}
