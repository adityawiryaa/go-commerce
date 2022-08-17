package product

import (
	"context"
	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/core/domain/entity"
	"go-commerce/services/product/core/domain/port"
)

func (productUC productUseCase) GetListProduct(ctx context.Context, payload *productPb.GetListProductRequest) ([]*entity.Product, uint32, error) {
	products, err := productUC.productRepo.FindProducts(ctx, port.FindOptions{
		IsPaginate: true,
		Limit:      int(payload.GetLimit()),
		Offset:     int(payload.GetOffset()),
	})
	if err != nil {
		return nil, 0, err
	}
	count, err := productUC.productRepo.CountProducts(ctx, port.FindOptions{
		IsPaginate: false,
	})
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}
