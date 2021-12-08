package main

import (
	"fmt"
	"github.com/qudj/fly_starling_rpc/config"
	"github.com/qudj/fly_starling_rpc/middlewares"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50054"
)

func main() {
	config.InitConfig()

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	// 实例化grpc Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(middlewares.UnaryServerInterceptor),
	)

	// 注册HelloService
	fccService := NewStarlingServiceServer()
	servbp.RegisterStarlingServiceServer(s, fccService)

	log.Println(fmt.Sprintf("listen to %s:", Address))
	s.Serve(listen)
}
