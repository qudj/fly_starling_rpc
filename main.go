package main

import (
	"fmt"
	"github.com/qudj/fly_starling_rpc/config"
	servbp "github.com/qudj/fly_starling_rpc/models/fly_starling_serv"
	"google.golang.org/grpc"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50053"
)

func main() {
	config.InitConfig()

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	fccService := NewStarlingServiceServer()
	servbp.RegisterStarlingServiceServer(s, fccService)

	fmt.Println(fmt.Sprintf("listen to %s:", Address))
	s.Serve(listen)
}
