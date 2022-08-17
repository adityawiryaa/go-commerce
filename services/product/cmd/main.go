package main

import (
	"context"
	"database/sql"
	"go-commerce/services/product/infrastructure/config"
	"go-commerce/services/product/infrastructure/grpc_server/handler"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

var (
	cfg    = config.InitProductConfig()
	ctx, _ = context.WithCancel(context.Background())
	db     = config.InitMysqlConfig()
)

func main() {
	defer func(db *sql.DB) { _ = db.Close() }(db)

	/* Run Init RPC Server */
	go RunGrpcServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}
}

func RunGrpcServer() {
	grpcServer := grpc.NewServer()
	handler.NewHandlerProductService(grpcServer, db)

	lis, err := net.Listen("tcp", cfg.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	go func() {
		log.Println("Grpc Server listen to " + cfg.RpcPort)
		log.Fatal(grpcServer.Serve(lis))
	}()
}
