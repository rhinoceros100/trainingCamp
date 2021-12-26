package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//go build -o helloworld main.go wire_gen.go

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	//开启grpc服务器
	svr := InitSvr()

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

	//起svr监听
	lis, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		return
	}
	if err := svr.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
		return
	}
}
