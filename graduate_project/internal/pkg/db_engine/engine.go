package db_engine

import (
	"comm/config"
	"comm/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
	"time"
)

var Engine *xorm.Engine

func Init() error {
	log.Debug("db::data_source :", config.Config.DefaultString("db::data_source", ""))
	log.Debug("db::max_idle_conns :", config.Config.DefaultInt("db::max_idle_conns", 0))
	log.Debug("db::max_open_conns :", config.Config.DefaultInt("db::max_open_conns", 0))
	log.Debug("db::max_cache_num :", config.Config.DefaultInt("db::max_cache_num", 0))

	var err error
	Engine, err = xorm.NewEngine("mysql", config.Config.DefaultString("db::data_source", ""))
	Engine.SetMaxIdleConns(config.Config.DefaultInt("db::max_idle_conns", 100))
	Engine.SetMaxOpenConns(config.Config.DefaultInt("db::max_open_conns", 300))
	max_cache_num := config.Config.DefaultInt("db::max_cache_num", 0)
	if max_cache_num != 0 {
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), max_cache_num)
		Engine.SetDefaultCacher(cacher)
	}

	logFileName := config.Config.DefaultString("db::log_file", "./log/db.log")
	dbLogFile, _ := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
	if dbLogFile != nil {
		logger := xorm.NewSimpleLogger(dbLogFile)
		logger.SetLevel(core.LOG_WARNING)
		Engine.SetLogger(logger)
		Engine.ShowSQL(true)
		Engine.ShowExecTime(true)
	}
	go func() {
		for {
			<-time.After(time.Minute * 5)
			Engine.Ping()
		}
	}()
	return err
}
