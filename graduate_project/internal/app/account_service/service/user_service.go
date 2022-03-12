package service

import (
	"account_service/biz"
	"account_service/data"
	"context"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/account"
)

type UserService struct{}

func (us *UserService) GetAccount(ctx context.Context, in *account.GetAccountReq) (*account.GetAccountReply, error) {
	reply := &account.GetAccountReply{
		//ErrorCode: error_code.ERR_OK,
		ErrorCode: 0,
	}

	uid := in.GetUid()
	userRepo := &data.UserRepo{}
	userBiz := &biz.UserBiz{
		Data: userRepo,
	}
	userInfo, err := userBiz.GetAccount(ctx, uid)

	if nil != err {
		//reply.ErrorCode = error_code.ERR_SVR_INTERVAL
		reply.ErrorCode = 1001
		return reply, err
	}

	reply.UserInfo = &account.User{
		Name: userInfo.Name,
		Age:  userInfo.Age,
	}
	return reply, nil
}
