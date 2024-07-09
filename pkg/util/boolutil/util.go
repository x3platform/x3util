package boolutil

import (
	"strings"
)

// 转换支持的文本格式为 bool 类型
func Bool(text string) bool {
	text = strings.ToUpper(text)
	// 支持的格式
	// 支持 1 或 0
	// 支持 True 或 False
	// 支持 Yes 或 No
	// 支持 On 或 Off
	if text == "1" || text == "TRUE" || text == "YES" || text == "ON" {
		return true
	}

	return false
}
