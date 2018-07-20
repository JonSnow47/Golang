package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func DateString() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%v%.2v%.2v", y, int(m), d)
}

func RandInt() int {
	// 生成1000-9999之间任意数字
	return 1000 + rand.Intn(8999)
}

var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/~!@#$%^&*()_="
	codeLen = len(codes)
)

func RandString(len int) string {


	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}
