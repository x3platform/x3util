package randomutil

import (
	"math/rand"
	"strings"
	"time"
	"unsafe"

	"github.com/google/uuid"
)

// UUID 字符
func UUID() string {
	return uuid.NewString()
}

func UUID32() string {
	return strings.Replace(UUID(), "-", "", -1)
}

func UUID16() string {
	return UUID32()[:16]
}

// --------------------------------------------------------
// 生成随机字符
// --------------------------------------------------------

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func Chars(chars string, length int) string {
	b := make([]byte, length)

	// sleep 1ns 避免获取相同的时间生成相同的随机因子
	time.Sleep(1 * time.Nanosecond)
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))

	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := length-1, randomNumber.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = randomNumber.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(chars) {
			b[i] = chars[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	// 使用 unsafe 包来避免最终的拷贝
	return *(*string)(unsafe.Pointer(&b))
}

const randomNumbers = "1234567890"

// 随机数字
func Number(length int) string {
	return Chars(randomChars, length)
}

const randomChars = "abcdefghijklmnopqrstuvwxyz1234567890"

// 随机字符
func Characters(length int) string {
	return Chars(randomChars, length)
}

const randomComplexChars = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 复杂字符
func ComplexCharacters(length int) string {
	return Chars(randomComplexChars, length)
}
