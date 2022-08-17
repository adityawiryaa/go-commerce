package main

import (
	"context"
	"log"
	"net/http"

	"go-commerce/services/api/infrastructure/config"
	"go-commerce/services/api/infrastructure/handler"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cfg = config.InitApiConfig()
)

func main() {
	gw := runtime.NewServeMux()
	RunGrpcServer(gw)
	gwServer := &http.Server{
		Addr:    cfg.AppPort,
		Handler: gw,
	}
	log.Println("Serving gRPC-Gateway on " + cfg.AppPort)
	log.Fatalln(gwServer.ListenAndServe())
}

func RunGrpcServer(gw *runtime.ServeMux) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	/* HANDLER Service */
	go handler.NewUserHandler(context.Background(), gw, opts, cfg)
	go handler.NewProductHandler(context.Background(), gw, opts, cfg)
}
