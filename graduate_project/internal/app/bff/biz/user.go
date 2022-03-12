package biz

import (
	"bff/data"
	"context"
	"errors"
)

type UserBiz struct {
	data *data.UserRepo
}

func NewUserBiz(repo *data.UserRepo) *UserBiz {
	return &UserBiz{
		data: repo,
	}
}

func (ub *UserBiz) GetUserInfo(ctx context.Context, uid uint64) (*data.UserInfo, error) {
	if uid == 0 {
		return nil, errors.New("invalid user id")
	}

	userInfo, err := ub.data.GetUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
