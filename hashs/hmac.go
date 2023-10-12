package hashs

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"hash/fnv"
	"io"
	"time"
)

func hmacHash(msg string, key string) (hashData []byte) {
	k := []byte(key)
	mac := hmac.New(sha1.New, k)
	io.WriteString(mac, msg)
	return mac.Sum(nil)
}

func fnvAlgo(msg string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(msg))
	return h.Sum64()
}

func TestifyHmac() {
	start := time.Now()
	ip := "fdbd:dc61:ff::231:112:127:9561"
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("%d", i)
		sum := hmacHash(ip, key)
		fmt.Printf("%s, %x\n", ip, sum)
	}
	fmt.Println(time.Since(start))
}

func TestifyFnvAlgo() {
	start := time.Now()
	ip := "fdbd:dc61:ff::231:112:127:9561"
	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("%s:%d", ip, i)
		sum := fnvAlgo(msg)
		fmt.Printf("%s, %d\n", msg, sum)
	}
	fmt.Println(time.Since(start))
}
