package grpc_product

import (
	"context"

	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/infrastructure/grpc_server/response"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service GrpcProductService) GetListProduct(ctx context.Context, payload *productPb.GetListProductRequest) (*productPb.GetListProductResponse, error) {
	products, count, err := service.ProductUseCase.GetListProduct(ctx, payload)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return response.NewListProductResponse(products, count), nil
}
