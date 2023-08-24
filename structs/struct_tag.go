package structs

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"-"`
}

type Teacher struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Courses []string `json:"courses,omitempty"`
}

// 测试结构体json "-"标签的作用
func TestifyIngoreLLabel() {
	tea := Teacher{
		Name:    "alan",
		Age:     18,
		Courses: []string{},
	}

	bytes, err := json.Marshal(&tea)
	if err != nil {
		panic("Marshal failed")
	}

	fmt.Println(string(bytes)) // {"name":"alan","age":18}
}
