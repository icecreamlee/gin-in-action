package configs

import (
	"github.com/IcecreamLee/goutils"
	"gopkg.in/ini.v1"
)

var (
	ENV  = "dev"
	Port = "8000"
)

func init() {
	RootPath := goutils.GetCurrentPath()
	iniConfig, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, RootPath+"app.ini")
	if err != nil {
		panic("Load config file failure: " + err.Error())
	}
	section := iniConfig.Section("app")
	ENV = section.Key("env").String()
	Port = section.Key("port").String()
	section = iniConfig.Section("db")
	DBHost = section.Key("host").String()
	DBPort = section.Key("port").String()
	DBUser = section.Key("user").String()
	DBPassword = section.Key("password").String()
	DBName = section.Key("db_name").String()
	section = iniConfig.Section("redis")
	RedisAddr = section.Key("addr").String()
	RedisPassword = section.Key("password").String()
	RedisDB, _ = section.Key("db").Int()
}
