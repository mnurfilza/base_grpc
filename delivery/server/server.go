package server

import (
	"base_grpc/delivery/grpc/car"
	ucCar "base_grpc/usecase/car"

	"base_grpc/delivery/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port = 8080

func RunServer() {

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			unaryAuthInterceptor,
			unaryTraceInterceptor,
		),
	}
	srv := grpc.NewServer(
		opts...,
	)
	reflection.Register(srv)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	proto.RegisterCarObjectServer(srv, car.NewCarSer(ucCar.NewCar()))

	srv.Serve(lis)

}
