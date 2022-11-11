package pkg_http_server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	handlers "go-clean-architecture/todo/delivery/http"
	"go-clean-architecture/utils"
	response "go-clean-architecture/utils/response"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type HTTPServer interface {
	PrintAllRoutes()
	Run() error
	GracefulStop(ctx context.Context) error
	GetRouter() *chi.Mux
}

type HTTPServerImpl struct {
	router *chi.Mux
	svr    *http.Server
}

func NewHTTPServer(todoHandler handlers.TodoHTTPHandler) HTTPServer {
	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	router := chi.NewRouter()
	router.Use(
		sentryHandler.Handle,
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		// middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, response.H{
			"success": "true",
			"code":    200,
			"message": "Services run properly",
		})
	})

	// Register TodoHTTPHandler routes
	todoHandler.RegisterRoutes(router)

	return &HTTPServerImpl{
		router: router,
	}
}

// PrintAllRoutes - Walk and print out all routes
func (s *HTTPServerImpl) PrintAllRoutes() {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		logrus.Printf("%s %s\n", method, route)
		return nil
	}
	router := s.GetRouter()
	if err := chi.Walk(router, walkFunc); err != nil {
		utils.CaptureError(err)
	}
}

// Run - running server
func (s *HTTPServerImpl) Run() error {
	addr := fmt.Sprintf("%s%s", ":", os.Getenv("PORT"))
	logrus.Infoln("HTTP server listening on", addr)

	s.PrintAllRoutes()

	router := s.GetRouter()
	s.svr = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	err := s.svr.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// GracefulStop the server
func (s *HTTPServerImpl) GracefulStop(ctx context.Context) error {
	return s.svr.Shutdown(ctx)
}

func (s *HTTPServerImpl) GetRouter() *chi.Mux {
	return s.router
}
