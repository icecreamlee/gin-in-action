package models

import (
	"gin_in_action/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", configs.DBUser+":"+configs.DBPassword+"@tcp("+configs.DBHost+":"+configs.DBPort+")/"+configs.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}
