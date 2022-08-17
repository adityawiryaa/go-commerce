package handler

import (
	"database/sql"

	"google.golang.org/grpc"

	productPb "go-commerce/proto/_generated/product"
	"go-commerce/services/product/core/usecase/product"
	"go-commerce/services/product/infrastructure/grpc_server/service/grpc_product"
	"go-commerce/services/product/infrastructure/repository/mysql"
)

func NewHandlerProductService(grpcServer *grpc.Server, db *sql.DB) {
	productRepo := mysql.NewProductRepository(db)

	productPb.RegisterProductServiceServer(grpcServer, grpc_product.GrpcProductService{
		ProductUseCase: product.NewProductUseCase(
			productRepo,
		),
	})
}
