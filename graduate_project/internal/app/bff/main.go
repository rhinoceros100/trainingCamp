package main

import (
	"bff/service"
	"config"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Init() error {
	//初始化系统端口配置
	svrConf := "../../../configs/bff.json"
	err := config.Init(svrConf)
	if err != nil {
		fmt.Println("svr port config init err", err)
		return err
	}

	return nil
}

func httpSvr(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-ctx.Done()
		fmt.Println(addr, "Shutdown")
		s.Shutdown(ctx)
	}()
	fmt.Println("httpSvr", addr)
	return s.ListenAndServe()
}

func main() {
	//初始化日志和配置
	if Init() != nil {
		return
	}
	fmt.Println("main start")

	//http.HandleFunc("/query/user_info", (&service.HandleQuery{}).ServeHTTP)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	//开启http服务器
	eg.Go(func() error {
		httpPort := config.Config.DefaultString("ip_port", "127.0.0.1:22222")
		return httpSvr(ctx, httpPort, &service.HandleQuery{})
	})

	//信号注册
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	c := make(chan os.Signal, 1)
	signal.Notify(c, sigs...)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done")
				return ctx.Err()
			case <-c:
				fmt.Println("c cancel")
				cancel()
				return ctx.Err()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println(err.Error())
		return
	}

	httpPort := config.Config.DefaultString("port", ":22222")
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("failed to serve: %v", err)
		return
	}
}
