package repository

import (
	"context"
	"go-commerce/services/product/core/domain/entity"
	"go-commerce/services/product/core/domain/port"
)

type ProductRepositoryPort interface {
	AddProduct(ctx context.Context, product *entity.Product) error
	FindProduct(ctx context.Context, payload port.FindOptions) (*entity.Product, error)
	FindProducts(ctx context.Context, payload port.FindOptions) ([]*entity.Product, error)
	CountProducts(ctx context.Context, payload port.FindOptions) (uint32, error)
}
