package service

import (
	"bff/biz"
	"bff/data"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HandleQuery struct {
	biz *biz.UserBiz
}

func (hq *HandleQuery) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("ServeHTTP")
	_ = req.ParseForm()
	uid := hq.ParseUint64(req, "uid")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	//获取玩家数据
	userRepo := &data.UserRepo{}
	userBiz := biz.NewUserBiz(userRepo)
	hq.biz = userBiz
	userInfo, err := hq.biz.GetUserInfo(ctx, uid)
	if nil != err {
		ui := data.NewUserInfo(uid)
		_, _ = resp.Write(hq.Marshal(ui))
		return
	}

	_, _ = resp.Write(hq.Marshal(userInfo))
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
	jsonRet, err := json.Marshal(s)
	if err != nil {
		return []byte("")
	}
	return jsonRet
}
