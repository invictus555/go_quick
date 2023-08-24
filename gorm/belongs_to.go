package gorm

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// `Employee` belongs to `Company`，`CompanyID` 是外键
type Employee struct {
	gorm.Model
	CompanyRefer int
	Name         string
	Company      Company `gorm:"foreignKey:CompanyRefer"` // 指定CompanyRefer为外键，关联主表Company的ID
}

type Company struct {
	ID   int
	Name string
}

func TestifyBelongsTo() {
	// Employee里面有 company表的结构 所以只需要自动迁移Employee表即
	if err := db.AutoMigrate(&Employee{}); err != nil {
		panic(err)
	}

	var finds []Employee
	employee := &Employee{
		Name: "shengchao",
		Company: Company{
			ID:   1234,
			Name: "FeiShu",
		},
	}
	condition := &Employee{
		Name:         employee.Name,
		CompanyRefer: employee.Company.ID,
	}

	tx := db.Model(&Employee{}).Where(condition).Preload("Company").First(&finds)
	if tx.Error == nil && len(finds) > 0 {
		fmt.Printf("have found %d records, err=%v\n", len(finds), tx.Error)
		for _, v := range finds {
			body, _ := json.Marshal(&v)
			fmt.Println(string(body)) // 打印输出查询结果
			if v.Name == condition.Name && v.CompanyRefer == condition.CompanyRefer {
				fmt.Println("this employee has been enrolled")
				return
			}
		}
		return // 能找记录则不必再创建
	}

	if err := db.Create(employee).Error; err != nil {
		panic(err)
	}
}
