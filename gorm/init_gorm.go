package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	// 数据库名:ker_service, 用户名:root, 密码:123456
	dsn = "root:123456@tcp(127.0.0.1:3306)/ker_service?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = instance
}
