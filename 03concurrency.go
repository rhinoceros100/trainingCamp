package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	//开启http服务器
	eg.Go(func() error {
		return httpSvr(ctx, "127.0.0.1:12345", &MockHandler{})
	})
	eg.Go(func() error {
		return httpSvr(ctx, "127.0.0.1:33333", &MockHandler{})
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
}

//mock http handler
type MockHandler struct {}

func (m *MockHandler) ServeHTTP(http.ResponseWriter, *http.Request)  {
	fmt.Println("MockHandler ServeHTTP")
}

func httpSvr(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server {
		Addr: addr,
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
