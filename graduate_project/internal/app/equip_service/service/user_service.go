package service

import (
	"context"
	"equip_service/biz"
	"equip_service/data"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/equip"
)

type UserService struct{}

func (us *UserService) GetEquip(ctx context.Context, in *equip.GetEquipReq) (*equip.GetEquipReply, error) {
	reply := &equip.GetEquipReply{
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

	reply.EquipInfo = &equip.Equip{
		BoxNum:  userInfo.BoxNum,
		BootNum: userInfo.BootNum,
	}
	return reply, nil
}
