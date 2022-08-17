package product

import (
	"go-commerce/services/product/core/domain/repository"
	"go-commerce/services/product/core/domain/usecase"
)

type productUseCase struct {
	productRepo repository.ProductRepositoryPort
}

func NewProductUseCase(productRepo repository.ProductRepositoryPort) usecase.ProductUseCasePort {
	return productUseCase{
		productRepo: productRepo,
	}
}
