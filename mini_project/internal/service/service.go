package service

import (
	"context"

	"github.com/go-redis/redis/v7"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain/dto"
	"github.com/wiryawanyurih/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, productDTO dto.CreateProductDTO) (domain.Product, error)
	Update(ctx context.Context, productDTO dto.UpdateProductDTO,
		productID primitive.ObjectID) (domain.Product, error)
}

type Carts interface {
	FindAll(ctx context.Context) ([]domain.Cart, error)
	FindByID(ctx context.Context, cartID primitive.ObjectID) (domain.Cart, error)
	FindCartItems(ctx context.Context, cartID primitive.ObjectID) ([]domain.CartItem, error)
	AddCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error)
	ClearCart(ctx context.Context, cartID primitive.ObjectID) error
	Create(ctx context.Context, cartDTO dto.CreateCartDTO) (domain.Cart, error)
	Update(ctx context.Context, cartDTO dto.UpdateCartDTO,
		cartID primitive.ObjectID) (domain.Cart, error)
	Delete(ctx context.Context, cartID primitive.ObjectID) error
}

type Orders interface {
	FindAll(ctx context.Context) ([]domain.Order, error)
	FindByID(ctx context.Context, orderID primitive.ObjectID) (domain.Order, error)
	Create(ctx context.Context, orderDTO dto.CreateOrderDTO) (domain.Order, error)
	Update(ctx context.Context, orderDTO dto.UpdateOrderDTO,
		orderID primitive.ObjectID) (domain.Order, error)
	Delete(ctx context.Context, orderID primitive.ObjectID) error
}

type Services struct {
	Products Products
	Carts    Carts
	Orders   Orders
}

type Deps struct {
	Repos       *repository.Repositories
	Services    *Services
	RedisClient *redis.Client
}

func NewServices(deps Deps) *Services {
	productsService := NewProductsService(deps.Repos.Products)
	cartsService := NewCartsService(deps.Repos.Carts, productsService)
	ordersService := NewOrdersService(deps.Repos.Orders, productsService, cartsService)

	return &Services{
		Products: productsService,
		Carts:    cartsService,
		Orders:   ordersService,
	}
}
