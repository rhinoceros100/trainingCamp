package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/account"
	"time"
)

type Service struct{}

func (svc *Service) GetAccount(ctx context.Context, in *account.GetAccountReq) (*account.GetAccountReply, error) {
	reply := &account.GetAccountReply{
		//ErrorCode: error_code.ERR_OK,
		ErrorCode: 0,
	}
	var err error = nil
	redisClient := svc.getRedisClient()
	if nil == redisClient {
		//reply.ErrorCode = error_code.ERR_SVR_INTERVAL
		reply.ErrorCode = 1001
		err = errors.New("Account svc redis")
		return reply, err
	}

	uid := in.GetUid()
	key := fmt.Sprintf("U%d", uid)
	r := redisClient.Get(key)
	name := r.String()
	userInfo := &account.User{
		Name: name,
	}
	reply.UserInfo = userInfo
	return reply, nil
}

func (svc *Service) getRedisClient() *redis.Client {
	redisOptions := &redis.Options{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
		DialTimeout:        time.Duration(time.Second * 10),
		ReadTimeout:        time.Duration(time.Second * 30),
		WriteTimeout:       time.Duration(time.Second * 30),
		PoolSize:           15,
		PoolTimeout:        time.Duration(time.Second * 30),
		IdleTimeout:        time.Duration(time.Millisecond * 500),
		IdleCheckFrequency: time.Duration(time.Millisecond * 500),
		Password:           "",
	}

	redisClient := redis.NewClient(redisOptions)
	if nil == redisClient {
		fmt.Println("nil == client")
		return nil
	}

	return redisClient
}
