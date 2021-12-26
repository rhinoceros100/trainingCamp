//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"helloworld/internal/biz"
	"helloworld/internal/data"
	"helloworld/internal/server"
	"helloworld/internal/service"
)

func InitSvr() *grpc.Server {
	wire.Build(server.NewGRPCServer, service.NewHelloService, biz.NewHelloUsecase, data.NewHelloRepo)
	return &grpc.Server{}
}
