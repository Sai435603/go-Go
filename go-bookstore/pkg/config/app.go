package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbConfig *gorm.DB

func connectToDB() {
	d, err := gorm.Open("mysql", "sai:password123@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbConfig = d
}

func GetDB() *gorm.DB {
	if dbConfig == nil {
		connectToDB()
	}
	return dbConfig
}
