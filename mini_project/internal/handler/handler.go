package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wiryawanyurih/ecommerce-api/internal/config"
	v1 "github.com/wiryawanyurih/ecommerce-api/internal/handler/v1"
	"github.com/wiryawanyurih/ecommerce-api/internal/service"
)

type Handler struct {
	config   *config.Config
	services *service.Services
}

func NewHandler(services *service.Services, config *config.Config) *Handler {

	return &Handler{
		config:   config,
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.Default())

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.config)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
