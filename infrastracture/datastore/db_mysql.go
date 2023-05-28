package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB() *gorm.DB {
	DB_DSN := "root:root@tcp(localhost:3382)/oauth2_resource_develop?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DB_DSN), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	return db
}
