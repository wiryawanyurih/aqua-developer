package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wiryawanyurih/ecommerce-api/internal/config"
	"github.com/wiryawanyurih/ecommerce-api/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initProductsRoutes(v1)
		h.initCartRoutes(v1)
		h.initOrdersRoutes(v1)
	}
}

func getIdFromPath(c *gin.Context, paramName string) (primitive.ObjectID, error) {
	idString := c.Param(paramName)
	if idString == "" {
		return primitive.ObjectID{}, errors.New("empty id param")
	}

	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id param")
	}

	return id, nil
}

func getIdFromRequestContext(context *gin.Context, paramName string) (primitive.ObjectID, error) {
	idString, ok := context.Get(paramName)
	if !ok {
		return primitive.ObjectID{}, errors.New("not authenticated")
	}

	id, err := primitive.ObjectIDFromHex(idString.(string))
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id param")
	}

	return id, nil
}
