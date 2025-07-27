package main

import (
	"log"
	"net"

	handler "github.com/empaid/ecom/services/orders/handler/orders"
	"github.com/empaid/ecom/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService)

	log.Println("Starting gRPC server on", s.addr)
	return gRPCServer.Serve(lis)
}
