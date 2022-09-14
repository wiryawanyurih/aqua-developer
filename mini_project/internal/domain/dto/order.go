package dto

import (
	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
	"time"
)

type CreateOrderDTO struct {
	OrderItems  []domain.OrderItem `json:"-"`
	ContactInfo domain.ContactInfo `json:"contactInfo"`
}

type UpdateOrderDTO struct {
	DeliveredAt time.Time `json:"deliveredAt"`
	Status      string    `json:"status"`
}

type UpdateOrderInput struct {
	DeliveredAt time.Time `json:"deliveredAt"`
	Status      string    `json:"status"`
}
