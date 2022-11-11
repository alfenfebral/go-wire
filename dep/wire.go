//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package dep

import (
	pkg_logger "go-clean-architecture/pkg/logger"
	pkg_mongodb "go-clean-architecture/pkg/mongodb"
	pkg_server "go-clean-architecture/pkg/server"
	pkg_http_server "go-clean-architecture/pkg/server/http"
	handlers "go-clean-architecture/todo/delivery/http"
	repository "go-clean-architecture/todo/repository"
	service "go-clean-architecture/todo/service"

	"github.com/google/wire"
)

func InitializeServer() (pkg_server.Server, error) {
	wire.Build(
		pkg_logger.NewLogger,
		pkg_mongodb.NewMongoDB,
		repository.NewMongoTodoRepository,
		service.NewTodoService,
		pkg_http_server.NewHTTPServer,
		handlers.NewTodoHTTPHandler,
		pkg_server.NewServer,
	)

	return &pkg_server.ServerImpl{}, nil
}
