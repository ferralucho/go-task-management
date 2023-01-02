// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/ferralucho/go-task-management/internal/usecase/repo/trello-api"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/ferralucho/go-task-management/config"
	v1 "github.com/ferralucho/go-task-management/internal/controller/http/v1"
	"github.com/ferralucho/go-task-management/internal/usecase"
	"github.com/ferralucho/go-task-management/pkg/httpserver"
	"github.com/ferralucho/go-task-management/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Use case
	cardUseCase := usecase.New(
		trello_api.New(),
	)

	var err error

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, cardUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
