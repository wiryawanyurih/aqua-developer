package v1

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", h.getAllProducts)
		products.GET("/:id", h.getProductById)

	}
}

// GetProducts
func (h *Handler) getAllProducts(context *gin.Context) {
	products, err := h.services.Products.FindAll(context.Request.Context())
	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	productsArray := make([]domain.Product, len(products))
	if products != nil {
		productsArray = products
	}

	time.Sleep(5 * time.Second)

	successResponse(context, productsArray)
}

// GetProductById
func (h *Handler) getProductById(context *gin.Context) {
	id, err := getIdFromPath(context, "id")
	if err != nil {
		badRequestResponse(context,
			fmt.Sprintf("invalid product id param"), err)
		return
	}

	product, err := h.services.Products.FindByID(context.Request.Context(), id)
	if err != nil {
		notFoundOrInternalErrorResponse(context,
			fmt.Sprintf("no products with id: %s", id.Hex()), err)
		return
	}

	successResponse(context, product)
}
