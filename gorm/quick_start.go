package gorm

import (
	"fmt"
)

func QuickStartDemo() {
	// 迁移 schema并创建table
	if err := db.AutoMigrate(&DBAppInfo{}); err != nil {
		panic("AutoMigrate failed: info = " + err.Error())
	}

	var rets []DBAppInfo
	condition := &DBAppInfo{
		AppID:   1233,
		AppName: "TikTok",
	}

	tx := db.Where(condition).Find(&rets)
	if len(rets) > 0 || tx == nil {
		fmt.Println("find record, return it", rets)
		return
	}

	appInfo := &DBAppInfo{
		AppID:   1233,
		AppName: "TikTok",
		Creator: "shengchao",
		Admins:  "shengchao;gaohonglei",
	}
	tx = tx.Create(appInfo) // Create a record
	if tx.Error != nil {
		fmt.Println("Create return error:", tx.Error)
	} else {
		fmt.Printf("RowsAffected: %d, PrimaryKey: %v", tx.RowsAffected, appInfo.ID)
	}
}
