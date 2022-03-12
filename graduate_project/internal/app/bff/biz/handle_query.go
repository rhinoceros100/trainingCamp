package biz

import (
	"config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/account"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/equip"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type HandleQuery struct {
}

func (hq *HandleQuery) QueryUser(resp http.ResponseWriter, req *http.Request) {
	uid := hq.ParseUint64(req, "uid")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	userInfo := &UserInfo{}

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done")
				return ctx.Err()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println(err.Error())
		return
	}

	accountSvrClient := hq.GetAccountSvrClient(ctx)
	equipSvrClient := hq.GetEquipSvrClient(ctx)
	if nil == accountSvrClient || nil == equipSvrClient {
		fmt.Println("nil == accountSvrClient || nil == equipSvrClient")
		return
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
	resp.Write(hq.Marshal(userInfo))
}

func (hq *HandleQuery) parseInt(req *http.Request, param string) int {
	if param == "" {
		return 0
	}

	param_arr, param_ok := req.Form[param]
	if !param_ok || param_arr[0] == "" {
		return 0
	}

	param_int, err := strconv.Atoi(param_arr[0])
	if nil != err {
		fmt.Println("pass para err, param:", param, "err:", err)
		return 0
	}
	return param_int
}

func (hq *HandleQuery) ParseUint64(req *http.Request, param string) uint64 {
	return uint64(hq.parseInt(req, param))
}

func (hq *HandleQuery) Marshal(s interface{}) []byte {
	json_ret, error := json.Marshal(s)
	if error != nil {
		return []byte("")
	}
	return json_ret
}

func (hq *HandleQuery) GetAccountSvrClient(ctx context.Context) account.AccountServiceClient {
	accountSvrAddr := config.Config.DefaultString("svr_addr::account_svr_addr", "127.0.0.1:12345")
	accountSvrConn, err := grpc.Dial(accountSvrAddr, grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", addr)
	}))
	if err != nil {
		return nil
	}
	client := account.NewAccountServiceClient(accountSvrConn)
	return client
}

func (hq *HandleQuery) GetEquipSvrClient(ctx context.Context) equip.EquipServiceClient {
	equipSvrAddr := config.Config.DefaultString("svr_addr::equip_svr_addr", "127.0.0.1:23456")
	equipSvrConn, err := grpc.Dial(equipSvrAddr, grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", addr)
	}))
	if err != nil {
		return nil
	}
	client := equip.NewEquipServiceClient(equipSvrConn)
	return client
}

type UserInfo struct {
	Uid     uint64
	Name    string
	Age     int32
	BoxNum  int32
	BootNum int32
}
