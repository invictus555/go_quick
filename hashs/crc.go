package hashs

import (
	"fmt"
	"hash/crc32"
	"time"
)

func TestifyCRC32() {
	start := time.Now()
	ip := "fdbd:dc61:ff::231:112:127:9561"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%s:%d", ip, i)
		sum := crc32.ChecksumIEEE([]byte(key))
		fmt.Println("%s, %d", key, sum)
	}
	fmt.Println(time.Since(start))
}
