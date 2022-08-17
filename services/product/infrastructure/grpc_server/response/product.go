package response

import (
	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/core/domain/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewProductResponse(product *entity.Product) *productPb.Product {
	return &productPb.Product{
		Id:        product.Id,
		Name:      product.Name,
		Quantity:  product.Quantity,
		Price:     product.Price,
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: timestamppb.New(product.UpdatedAt),
	}
}

func NewListProductResponse(products []*entity.Product, count uint32) *productPb.GetListProductResponse {
	responses := make([]*productPb.Product, 0)
	for _, product := range products {
		responses = append(responses, NewProductResponse(product))
	}

	return &productPb.GetListProductResponse{
		Items: responses,
		Count: count,
	}
}
