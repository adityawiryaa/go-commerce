package product

import (
	"context"
	"errors"
	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/core/domain/entity"
	"go-commerce/services/product/core/domain/port"
)

func (productUC productUseCase) CreateProduct(ctx context.Context, payload *productPb.CreateProductRequest) (*entity.Product, error) {
	isExist, _ := productUC.productRepo.FindProduct(ctx, port.FindOptions{
		Name: payload.GetName(),
	})
	if isExist != nil {
		return nil, errors.New("product name is already exists")
	}

	product := entity.NewProduct(entity.ProductDTO{
		Name:     payload.GetName(),
		Quantity: payload.GetQuantity(),
		Price:    payload.GetPrice(),
	})

	err := productUC.productRepo.AddProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
