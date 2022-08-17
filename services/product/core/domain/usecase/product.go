package usecase

import (
	"context"
	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/core/domain/entity"
)

type ProductUseCasePort interface {
	CreateProduct(ctx context.Context, payload *productPb.CreateProductRequest) (*entity.Product, error)
	GetListProduct(ctx context.Context, payload *productPb.GetListProductRequest) ([]*entity.Product, uint32, error)
}
