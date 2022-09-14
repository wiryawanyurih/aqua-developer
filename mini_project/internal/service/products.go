package service

import (
	"context"

	"github.com/wiryawanyurih/ecommerce-api/internal/domain"
	"github.com/wiryawanyurih/ecommerce-api/internal/domain/dto"
	"github.com/wiryawanyurih/ecommerce-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService struct {
	repo repository.Products
}

func (p *ProductsService) FindAll(ctx context.Context) ([]domain.Product, error) {
	products, err := p.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductsService) FindByID(ctx context.Context, productID primitive.ObjectID) (domain.Product, error) {
	product, err := p.repo.FindByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}
	return product, err
}

func (p *ProductsService) Create(ctx context.Context, product dto.CreateProductDTO) (domain.Product, error) {
	return p.repo.Create(ctx, domain.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Categories:  product.Categories,
	})
}

func (p *ProductsService) Update(ctx context.Context, productDTO dto.UpdateProductDTO, productID primitive.ObjectID) (domain.Product, error) {
	return p.repo.Update(ctx, dto.UpdateProductInput{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Categories:  productDTO.Categories,
	}, productID)
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{
		repo: repo,
	}
}
