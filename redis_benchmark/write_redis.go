package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	SLEN int = 100
	NUM  int = 100000
)

func getRandomString(len int) string {
	s := ""
	for i := 0; i < len; i++ {
		s += "a"
	}
	return s
}

func main() {
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
		return
	}

	s := getRandomString(SLEN)
	for i := 0; i < NUM; i++ {
		k := fmt.Sprintf("%d", i)
		sc := redisClient.Set(k, s, time.Duration(time.Second*100))
		err := sc.Err()
		if err != nil {
			fmt.Println("set redis err:", err, k)
			return
		}
	}
}
