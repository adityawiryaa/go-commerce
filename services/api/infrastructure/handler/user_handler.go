package handler

import (
	"context"
	"fmt"
	"log"

	"go-commerce/services/api/infrastructure/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	userPb "go-commerce/proto/_generated/user"
)

func NewUserHandler(ctx context.Context, gw *runtime.ServeMux, opts []grpc.DialOption, cfg *config.ApiConfig) {
	err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, gw, cfg.ServicePort.User, opts)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to register service %s : ", serviceUserName), err)
	}
}
