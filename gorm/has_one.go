package gorm

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Customer 有一张 CreditCard，UserID 是外键
type Customer struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Age        uint8      `gorm:"column:age"`                             // 年龄
	UID        string     `gorm:"primaryKey;column:uid;type:varchar(18)"` // 身份证ID 18位, 指定非ID为主键时，请删除gorm.Model
	Name       string     `gorm:"column:name;type:varchar(32)"`           // 姓名
	Addr       string     `gorm:"column:addr;type:varchar(256)"`          // 居住地
	CreditCard CreditCard `gorm:"foreignKey:CustomerID;references:UID"`   // 指定CustomerID为外键，并关联Customer.UID
}

type CreditCard struct {
	gorm.Model
	Number           string `gorm:"column:number"`            // 信用卡号
	CustomerID       string `gorm:"column:customer_id"`       // 外键
	RegisterLocation string `gorm:"column:register_location"` // 注册地
}

func TestifyHasOnCreate() {
	// 注意与belongs to的建表区别
	if err := db.AutoMigrate(&Customer{}, &CreditCard{}); err != nil {
		panic(err)
	}

	var finds []Customer
	newUser := &Customer{
		UID:  "513701198909023510",
		Age:  23,
		Addr: "28-7-601.xinlongcheng.changping.beijing",
		Name: "shengchao",
		CreditCard: CreditCard{
			Number:           "20230812142536",
			RegisterLocation: "12th ICBC,beijing",
		},
	}
	condition := &Customer{
		UID: newUser.UID, // 根据用户身份证查重
	}
	tx := db.Model(&Customer{}).Where(condition).Preload("CreditCard").Find(&finds)
	if tx.Error == nil && len(finds) > 0 {
		fmt.Printf("have found %d records\n", len(finds))
		for _, v := range finds {
			body, _ := json.Marshal(v)
			fmt.Println(string(body))
		}
		return
	}

	if err := db.Create(newUser).Error; err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
}

func TestifyHasOneDelete() {
	cond := &Customer{
		UID: "513701198909023510",
	}
	tx := db.Select("Customer").Where(cond).Delete(&CreditCard{})
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func TestifyHasOneUpdate() {

}
