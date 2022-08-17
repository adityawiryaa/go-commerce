package handler

import (
	"context"
	"fmt"
	"log"

	"go-commerce/services/api/infrastructure/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	productPb "go-commerce/proto/_generated/product"
)

func NewProductHandler(ctx context.Context, gw *runtime.ServeMux, opts []grpc.DialOption, cfg *config.ApiConfig) {
	err := productPb.RegisterProductServiceHandlerFromEndpoint(ctx, gw, cfg.ServicePort.Product, opts)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to register service %s : ", serviceProductName), err)
	}
}
