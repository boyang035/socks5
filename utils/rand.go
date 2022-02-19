package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GetRandomString2 生成n个随机字符串，n为偶数
func GetRandomString2(n int) string {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}