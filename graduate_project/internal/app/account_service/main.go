package main

import (
	"account_service/biz"
	"config"
	"fmt"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func Init() error {
	//初始化系统端口配置
	svrConf := "../../../configs/account.json"
	err := config.Init(svrConf)
	if err != nil {
		fmt.Println("svr port config init err", err)
		return err
	}

	return nil
}

func main() {
	//初始化日志和配置
	if Init() != nil {
		return
	}
	fmt.Println("main start")

	grpcPort := config.Config.DefaultString("port", ":12345")
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		return
	}

	fmt.Println("grpc service")
	svr := grpc.NewServer()
	account.RegisterAccountServiceServer(svr, &biz.Service{})
	reflection.Register(svr)
	if err := svr.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
		return
	}
}
