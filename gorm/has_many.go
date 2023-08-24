package gorm

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	ID    string `gorm:"column:id;type:varchar(18);primaryKey"` // 主键
	Age   int    `gorm:"column:age"`
	Grade int    `gorm:"column:grade"`
	Name  string `gorm:"column:name;type:varchar(64)"`
	Books []Book `gorm:"foreignKey:StudentID;references:ID"` // 指定外键与引用,一定要大写
}

type Book struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;column:id"` // 主键
	Category  uint8  `gorm:"column:category"`
	Name      string `gorm:"column:name;type:varchar(12)"`
	ISBN      string `gorm:"column:isbn;type:varchar(36)"`
	StudentID string `gorm:"column:student_id;type:varchar(18)"`
}

func TestifyHasManyCreate() {
	// 建表
	if err := db.AutoMigrate(&Student{}, &Book{}); err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}

	var finds []Student
	stu := &Student{
		Age:   15,
		Name:  "xiaohui",
		Grade: 9,
		ID:    "123456789987654321",
		Books: []Book{
			{
				Name:     "yuwen",
				ISBN:     "4B3D074D-C117-4032-9CCC-F7F4C3F57B7C",
				Category: 12,
			},
			{
				Name:     "shuxue",
				ISBN:     "5A232271-7630-4CCC-BBD9-E4BC333733EF",
				Category: 36,
			},
		},
	}

	// 查重
	tx := db.Model(&Student{}).Preload("Books").Where("id=?", stu.ID).Find(&finds)
	if tx.Error == nil && len(finds) > 0 {
		fmt.Printf("have found %d records\n", len(finds))
		for _, v := range finds {
			body, _ := json.Marshal(v)
			fmt.Println(string(body))
		}
		return
	}

	if tx := db.Create(stu); tx.Error != nil {
		fmt.Println(tx.Error)
		panic(tx.Error)
	} else {
		fmt.Println("success:", tx.RowsAffected)
	}
}

// 更新Book数据，将所有的yuwen的category换成20
func TestifyHasManyUpdate() {
	var finds []Student
	tx := db.Model(&Student{}).Preload("Books").Find(&finds)
	if tx.Error != nil || len(finds) == 0 {
		fmt.Println("not found records, err = ", tx.Error)
		return
	}

	// 更新数据
	for i, v := range finds {
		body, _ := json.Marshal(v)
		fmt.Println(string(body))
		if v.Age == 15 { // 更新student.Age
			finds[i].Age = 14
		}
		for j, book := range v.Books {
			if book.Name == "yuwen" { // 更新yuwen的category
				v.Books[j].Category = 20
			}
		}
	}

	for _, v := range finds {
		// 更新student信息与book信息
		db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&v)
	}
}
