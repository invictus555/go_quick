package mredis

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Teacher struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Courses []string `json:"courses,omitempty"`
}

func (t *Teacher) MarshalBinary() (data []byte, err error) {
	fmt.Println("MarshalBinary")
	return json.Marshal(t)
}

func (t *Teacher) UnmarshalBinary(data []byte) error {
	fmt.Println("UnmarshalBinary")
	return json.Unmarshal(data, t)
}

// 创建redis客户端
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.37.74.224:6379", // redis地址
		Password: "",                  // 密码
		DB:       0,                   // 使用默认数据库
	})
	return client
}

func TesifyGoRedis() {
	// 创建客户端
	client := newClient()
	defer client.Close()

	teacher := &Teacher{
		Name:    "VVV",
		Age:     35,
		Courses: []string{"yuwen", "shuxue"},
	}
	// 设置key
	err := client.HSet("teacher", "7th", teacher).Err()
	if err != nil {
		panic(err)
	}

	// 获取key
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)
}
