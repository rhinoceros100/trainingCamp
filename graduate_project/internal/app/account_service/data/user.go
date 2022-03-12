package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type UserRepo struct{}

type UserInfo struct {
	Name string
	Age  int32
}

func (ur *UserRepo) GetUserInfo(ctx context.Context, uid uint64) (*UserInfo, error) {
	var err error = nil
	redisClient := ur.getRedisClient()
	if nil == redisClient {
		err = errors.New("Account svc redis")
		return nil, err
	}

	key := fmt.Sprintf("U%d", uid)
	r := redisClient.Get(key)
	name := r.String()
	age, _ := r.Int()
	userInfo := &UserInfo{
		Name: name,
		Age:  int32(age),
	}
	return userInfo, nil
}

func (ur *UserRepo) getRedisClient() *redis.Client {
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
