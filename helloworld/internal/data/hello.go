package data

import (
	"context"
	"errors"
	"fmt"
)

type Hello struct {
	ID  uint64
	Msg string
}

type HelloRepo struct {
	ID  uint64
	Msg string
}

func NewHelloRepo() *HelloRepo {
	return &HelloRepo{
		ID:  111,
		Msg: "repo",
	}
}

func (hr *HelloRepo) GetHello(ctx context.Context, id uint64) (*Hello, error) {
	return &Hello{
		ID:  123,
		Msg: "Hello world",
	}, nil
}

func (hr *HelloRepo) SendHello(ctx context.Context, hello *Hello) error {
	if nil == hello {
		return errors.New("Hello nil")
	}

	fmt.Println("SendHello from id: %d, msg: %s", hello.ID, hello.Msg)
	return nil
}
