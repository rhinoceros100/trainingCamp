package biz

import (
	account "../api/account"
	"context"
)

type Service struct{}

func (svc *Service) GetAccount(ctx context.Context, in *account.GetAccountReq) (*account.GetAccountReply, error) {
	return nil, nil
}
