// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ferralucho/go-task-management/internal/usecase"
	"github.com/ferralucho/go-task-management/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Management API
// @description Api for space management
// @version     1.0
// @host        localhost:8082
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Card) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newTaskRoutes(h, t, l)
	}
}
