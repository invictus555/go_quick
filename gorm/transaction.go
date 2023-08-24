package gorm

import (
	"gorm.io/gorm"
)

func TestifyTransaction() {
	// 创建product_info
	if err := db.AutoMigrate(&DBAppInfo{}); err != nil {
		panic("create app_info table failed")
	}

	if err := db.AutoMigrate(&DBModuleInfo{}); err != nil {
		panic("create module_info table failed")
	}

	// 自动事务：同时插入一条记录
	err := db.Transaction(func(tx *gorm.DB) error {
		ret := tx.Create(&DBAppInfo{
			AppID:   1128,
			AppName: "DouYin",
			Creator: "shengchao",
			Admins:  "shengchao",
		})
		if ret.Error != nil {
			return ret.Error
		}

		ret = tx.Create(&DBModuleInfo{
			AppID:      1128,
			ModuleName: "vod",
			Creator:    "shengchao",
		})
		if ret.Error != nil {
			return ret.Error
		}

		// time.Sleep(time.Second * 5) // 为了检查是否写入留足时间
		return nil
	})

	if err != nil {
		panic(err)
	}
}
