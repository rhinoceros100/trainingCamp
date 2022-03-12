package db_engine

import (
	"config"
	"errors"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Redis struct {
	redisAddress       string
	redisPwd           string
	db                 int
	dialTimeout        time.Duration
	readTimeout        time.Duration
	writeTimeout       time.Duration
	poolSize           int
	poolTimeout        time.Duration
	idleTimeout        time.Duration
	idleCheckFrequency time.Duration

	redisClient *redis.Client
}

var RedisEngine *Redis

func InitRedis() error {
	RedisEngine = &Redis{}
	RedisEngine.redisAddress = config.Config.String("redis::address")
	RedisEngine.redisPwd = config.Config.String("redis::pwd")
	RedisEngine.db = config.Config.DefaultInt("redis::db_num", 15)
	RedisEngine.dialTimeout = time.Duration(config.Config.DefaultInt("redis::dial_timeout", 10)) * time.Second
	RedisEngine.readTimeout = time.Duration(config.Config.DefaultInt("redis::read_timeout", 30)) * time.Second
	RedisEngine.writeTimeout = time.Duration(config.Config.DefaultInt("redis::write_timeout", 30)) * time.Second
	RedisEngine.poolSize = config.Config.DefaultInt("redis::pool_size", 15)
	RedisEngine.poolTimeout = time.Duration(config.Config.DefaultInt("redis::pool_timeout", 30)) * time.Second
	RedisEngine.idleTimeout = time.Duration(config.Config.DefaultInt("redis::idle_timeout", 500)) * time.Millisecond
	RedisEngine.idleCheckFrequency = time.Duration(config.Config.DefaultInt("redis::idle_check_frenquency", 500)) * time.Millisecond

	return RedisEngine.init()
}

func (this *Redis) init() error {
	//log.Debug("redis init")
	redisOptions := &redis.Options{
		Addr:               this.redisAddress,
		DB:                 this.db,
		DialTimeout:        this.dialTimeout,
		ReadTimeout:        this.readTimeout,
		WriteTimeout:       this.writeTimeout,
		PoolSize:           this.poolSize,
		PoolTimeout:        this.poolTimeout,
		IdleTimeout:        this.idleTimeout,
		IdleCheckFrequency: this.idleCheckFrequency,
		Password:           this.redisPwd,
	}

	this.redisClient = redis.NewClient(redisOptions)
	if nil == this.redisClient {
		log.Error("nil == client")
		return errors.New("Create redis client err")
	}
	return nil
}

func (this *Redis) GetClient() *redis.Client {
	return this.redisClient
}
