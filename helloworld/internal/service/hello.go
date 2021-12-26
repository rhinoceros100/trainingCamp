package service

import (
	"context"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/internal/data"
)

type HelloService struct {
	hello *biz.HelloUsecase
}

func NewHelloService(hello *biz.HelloUsecase) *HelloService {
	return &HelloService{
		hello: hello,
	}
}

func (hs *HelloService) SendHello(ctx context.Context, req *pb.SendHelloReq) (*pb.SendHelloRply, error) {
	err := hs.hello.Send(ctx, &data.Hello{
		ID:  req.Uid,
		Msg: req.Msg,
	})
	return &pb.SendHelloRply{}, err
}

func (hs *HelloService) GetHello(ctx context.Context, req *pb.GetHelloReq) (*pb.GetHelloRply, error) {
	hello, err := hs.hello.Get(ctx, req.GetUid())
	if nil != err {
		return nil, err
	}
	return &pb.GetHelloRply{Msg: hello.Msg}, nil
}
