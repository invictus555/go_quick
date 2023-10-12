package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	// 数据库名:gorm_demo, 用户名:root, 密码:123456
	dsn = "root:123456@tcp(10.37.74.224:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = instance
}

func GetDB() *gorm.DB {
	return db
}
