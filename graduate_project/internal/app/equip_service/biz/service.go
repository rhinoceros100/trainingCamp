package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rhinoceros100/trainingCamp/graduate_project/api/equip"
	"time"
)

type Service struct{}

func (svc *Service) GetEquip(ctx context.Context, in *equip.GetEquipReq) (*equip.GetEquipReply, error) {
	reply := &equip.GetEquipReply{
		//ErrorCode: error_code.ERR_OK,
		ErrorCode: 0,
	}
	var err error = nil
	redisClient := svc.getRedisClient()
	if nil == redisClient {
		//reply.ErrorCode = error_code.ERR_SVR_INTERVAL
		reply.ErrorCode = 1001
		err = errors.New("Equip svc redis")
		return reply, err
	}

	uid := in.GetUid()
	key := fmt.Sprintf("EquipBox%d", uid)
	r := redisClient.Get(key)
	num, _ := r.Int()
	equipInfo := &equip.Equip{
		BoxNum: int32(num),
	}
	reply.EquipInfo = equipInfo
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
