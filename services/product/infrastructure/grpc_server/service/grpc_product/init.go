package grpc_product

import "go-commerce/services/product/core/domain/usecase"

type GrpcProductService struct {
	ProductUseCase usecase.ProductUseCasePort
}
