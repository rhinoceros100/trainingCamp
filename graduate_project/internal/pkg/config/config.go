package config

import (
	"github.com/astaxie/beego/config"
)

var Config config.Configer

func Init(file string) error {
	var err error
	Config, err = config.NewConfig("json", file)
	return err
}
