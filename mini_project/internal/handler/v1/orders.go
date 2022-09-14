package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain/dto"
)

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup) {

}

// GerUserOrders
func (h *Handler) getUserOrders(context *gin.Context) {
	orders, err := h.services.Orders.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(context, orders)
}

// CreateOrder
func (h *Handler) createOrder(context *gin.Context) {
	cart, err := h.services.Carts.FindAll(context.Request.Context())
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if len(cart) == 0 {
		errorResponse(context, http.StatusBadRequest, "user cart is empty")
		return
	}

	if len(cart[0].CartItems) == 0 {
		errorResponse(context, http.StatusBadRequest, "user cart is empty")
		return
	}

	orderItems := make([]domain.OrderItem, len(cart[0].CartItems))
	for i, cartItem := range cart[0].CartItems {
		orderItems[i] = domain.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
	}

	var createOrderDTO dto.CreateOrderDTO
	err = context.BindJSON(&createOrderDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	order, err := h.services.Orders.Create(context.Request.Context(), dto.CreateOrderDTO{
		OrderItems:  orderItems,
		ContactInfo: createOrderDTO.ContactInfo,
	})

	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Carts.ClearCart(context.Request.Context(), cart[0].ID)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, "cart can't be cleared")
		return
	}

	successResponse(context, order)
}
