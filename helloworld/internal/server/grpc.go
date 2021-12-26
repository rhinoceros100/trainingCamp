package server

import (
	"google.golang.org/grpc"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/service"
)

func NewGRPCServer(hello *service.HelloService) *grpc.Server {
	svr := grpc.NewServer()
	pb.RegisterHelloServiceServer(svr, hello)
	return svr
}
