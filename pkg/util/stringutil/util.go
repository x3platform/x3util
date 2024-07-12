package stringutil

import (
	"bytes"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

// 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

const randomChars = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 随机字符
func Random(length int) string {
	return RandomByChars(randomChars, length)
}

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandomByChars(chars string, length int) string {
	b := make([]byte, length)

	randomNumber := rand.NewSource(time.Now().UnixNano())

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

// 内嵌 bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
