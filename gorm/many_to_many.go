package gorm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type Users struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func TestifyMany2Many() {
	db, err := gorm.Open(
		sqlite.Open("e-seller.db"),
		&gorm.Config{
			SkipDefaultTransaction: true, // 关闭默认事务
		})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Users{}) // 当使用 GORM 的 AutoMigrate 为 User 创建表时，GORM 会自动创建user_languages连接表
	if err != nil {
		fmt.Println(err)
	}

	english := Language{
		Name: "english",
	}

	chinese := Language{
		Name: "chinese",
	}

	sean := Users{
		Languages: []Language{english, chinese},
	}

	db.Create(&sean)
}
