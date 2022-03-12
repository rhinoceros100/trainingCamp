package main

import (
	"config"
	"fmt"
	"job/biz"
	"log"
	"net/http"
)

func Init() error {
	//初始化系统端口配置
	svrConf := "../../../configs/job.json"
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

	(&biz.RunJob{}).Run()

	httpPort := config.Config.DefaultString("port", ":33445")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("failed to serve: %v", err)
		return
	}
}
