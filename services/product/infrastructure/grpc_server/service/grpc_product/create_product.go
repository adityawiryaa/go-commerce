package grpc_product

import (
	"context"

	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/infrastructure/grpc_server/response"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service GrpcProductService) CreateProduct(ctx context.Context, payload *productPb.CreateProductRequest) (*productPb.Product, error) {
	product, err := service.ProductUseCase.CreateProduct(ctx, payload)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return response.NewProductResponse(product), nil
}
