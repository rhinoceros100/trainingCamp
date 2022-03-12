package data

import (
	"config"
	"context"
	"errors"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/account"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/equip"
	"google.golang.org/grpc"
	"sync"
)

type UserRepo struct{}

type UserInfo struct {
	Uid     uint64
	Name    string
	Age     int32
	BoxNum  int32
	BootNum int32
}

func NewUserInfo(uid uint64) *UserInfo {
	return &UserInfo{
		Uid: uid,
	}
}

func (ur *UserRepo) GetUserInfo(ctx context.Context, uid uint64) (*UserInfo, error) {
	userInfo := NewUserInfo(uid)

	accountSvrClient := ur.getAccountSvrClient(ctx)
	equipSvrClient := ur.getEquipSvrClient(ctx)
	if nil == accountSvrClient || nil == equipSvrClient {
		return userInfo, errors.New("Get remote service error")
	}

	//并发获取玩家的账号信息和装备信息
	var wg sync.WaitGroup
	var accountRsp *account.GetAccountReply = nil
	wg.Add(1)
	go func() {
		defer wg.Done()
		getAccountReq := &account.GetAccountReq{
			Uid: uid,
		}
		accountRsp, _ = accountSvrClient.GetAccount(ctx, getAccountReq)
		userInfo.Name = accountRsp.GetUserInfo().GetName()
		userInfo.Age = accountRsp.GetUserInfo().GetAge()
	}()

	wg.Add(1)
	var equipRsp *equip.GetEquipReply = nil
	go func() {
		defer wg.Done()
		getEquipReq := &equip.GetEquipReq{
			Uid: uid,
		}
		equipRsp, _ = equipSvrClient.GetEquip(ctx, getEquipReq)
		userInfo.BoxNum = equipRsp.GetEquipInfo().GetBoxNum()
		userInfo.BootNum = equipRsp.GetEquipInfo().GetBootNum()
	}()

	//获取信息后回包
	wg.Wait()
	return userInfo, nil
}

func (ur *UserRepo) getAccountSvrClient(ctx context.Context) account.AccountServiceClient {
	accountSvrAddr := config.Config.DefaultString("svr_addr::account_svr_addr", "127.0.0.1:12345")
	accountSvrConn, err := grpc.DialContext(ctx, accountSvrAddr)
	if err != nil {
		return nil
	}
	client := account.NewAccountServiceClient(accountSvrConn)
	return client
}

func (ur *UserRepo) getEquipSvrClient(ctx context.Context) equip.EquipServiceClient {
	equipSvrAddr := config.Config.DefaultString("svr_addr::equip_svr_addr", "127.0.0.1:23456")
	equipSvrConn, err := grpc.DialContext(ctx, equipSvrAddr)
	if err != nil {
		return nil
	}
	client := equip.NewEquipServiceClient(equipSvrConn)
	return client
}
