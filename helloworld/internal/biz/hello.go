package biz

import (
	"context"
	"helloworld/internal/data"
)

type HelloUsecase struct {
	repo *data.HelloRepo
}

func NewHelloUsecase(repo *data.HelloRepo) *HelloUsecase {
	return &HelloUsecase{repo: repo}
}

func (uc *HelloUsecase) Get(ctx context.Context, id uint64) (p *data.Hello, err error) {
	return uc.repo.GetHello(ctx, id)
}

func (uc *HelloUsecase) Send(ctx context.Context, hello *data.Hello) error {
	return uc.repo.SendHello(ctx, hello)
}
