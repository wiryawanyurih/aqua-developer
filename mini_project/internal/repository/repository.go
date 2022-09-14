package repository

import (
	"context"

	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Products interface {
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error)
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, productInput dto.UpdateProductInput,
		productID primitive.ObjectID) (domain.Product, error)
	Delete(ctx context.Context, productID primitive.ObjectID) error
}

type Carts interface {
	FindAll(ctx context.Context) ([]domain.Cart, error)
	FindByID(ctx context.Context, cartID primitive.ObjectID) (domain.Cart, error)
	FindCartItems(ctx context.Context, cartID primitive.ObjectID) ([]domain.CartItem, error)
	AddCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItem domain.CartItem, cartID primitive.ObjectID) (domain.CartItem, error)
	ClearCart(ctx context.Context, cartID primitive.ObjectID) error
	Create(ctx context.Context, cart domain.Cart) (domain.Cart, error)
	Update(ctx context.Context, cartInput dto.UpdateCartInput,
		cartID primitive.ObjectID) (domain.Cart, error)
	Delete(ctx context.Context, cartID primitive.ObjectID) error
}

type Orders interface {
	FindAll(ctx context.Context) ([]domain.Order, error)
	FindByID(ctx context.Context, orderID primitive.ObjectID) (domain.Order, error)
	Create(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, orderInput dto.UpdateOrderInput,
		orderID primitive.ObjectID) (domain.Order, error)
	Delete(ctx context.Context, orderID primitive.ObjectID) error
}

type Repositories struct {
	Products Products
	Carts    Carts
	Orders   Orders
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Products: NewProductsRepo(db),
		Carts:    NewCartsRepo(db),
		Orders:   NewOrdersRepo(db),
	}
}
