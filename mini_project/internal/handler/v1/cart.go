package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {
	cart := api.Group("/cart")
	{
		cart.POST("/", h.createCart)
	}

	cartItem := api.Group("/cartItem", h.extractCartId)
	{
		cartItem.GET("/", h.getCartItems)
		cartItem.POST("/", h.createCartItem)
		cartItem.DELETE("/", h.clearCart)
		cartItem.PUT("/:productID", h.updateCartItem)
	}
}

func (h *Handler) extractCartId(context *gin.Context) {
	newCart, err := h.services.Carts.Create(context.Request.Context(), dto.CreateCartDTO{
		ExpireAt: time.Now().Add(30 * time.Hour * 24),
	})
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.Set("cartID", newCart.ID)
}

// CreateCart
func (h *Handler) createCart(context *gin.Context) {
	newCart, err := h.services.Carts.Create(context.Request.Context(), dto.CreateCartDTO{
		ExpireAt: time.Now().Add(1 * time.Hour * 24),
	})
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	createdResponse(context, newCart.ID.Hex())
}

// GetCartItems
func (h *Handler) getCartItems(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	cartItems, err := h.services.Carts.FindCartItems(context, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, cartItems)
}

// CreateCartItem
func (h *Handler) createCartItem(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	var cartItemInput domain.CartItem
	err := context.BindJSON(&cartItemInput)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	cartItem, err := h.services.Carts.AddCartItem(context, cartItemInput, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	createdResponse(context, cartItem)
}

// UpdateCartItem
func (h *Handler) updateCartItem(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	productID, err := getIdFromPath(context, "productID")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	var cartItemInput dto.UpdateCartItemDTO
	err = context.BindJSON(&cartItemInput)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	cartItem, err := h.services.Carts.UpdateCartItem(context, domain.CartItem{
		ProductID: productID,
		Quantity:  cartItemInput.Quantity,
	}, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, cartItem)
}

// ClearCart
func (h *Handler) clearCart(context *gin.Context) {
	cartIDHex, ok := context.Get("cartID")
	if !ok {
		errorResponse(context, http.StatusInternalServerError, "failed to get cart id")
		return
	}
	cartID := cartIDHex.(primitive.ObjectID)

	err := h.services.Carts.ClearCart(context, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

// DeleteCart
func (h *Handler) deleteCartAdmin(context *gin.Context) {
	cartID, err := getIdFromPath(context, "id")
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Carts.Delete(context, cartID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}
