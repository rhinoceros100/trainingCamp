package biz

import (
	"account_service/data"
	"context"
	"errors"
)

type UserBiz struct {
	Data *data.UserRepo
}

func (ub *UserBiz) GetAccount(ctx context.Context, uid uint64) (*data.UserInfo, error) {
	if uid == 0 {
		return nil, errors.New("invalid user id")
	}

	userInfo, err := ub.Data.GetUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
