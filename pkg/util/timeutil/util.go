package timeutil

import (
	"strings"
	"time"
)

// 创建 time.Time 结构体
// @param  value  int64  时间搓
func Time(value int64) time.Time {
	return time.Unix(value, 0)
}

// 将 "yyyy-MM-dd HH:mm:ss" 转换为 Golang 时间格式 "2006-01-02 15:04:05"
func ConvetTimeLayout(layout string) string {
	// "yyyy-MM-dd HH:mm:ss" => "2006-01-02 15:04:05"
	// "dd/MMM/yyyy:HH:mm:ss Z" => "02/Jan/2006:15:04:05 -0700"
	replacer := strings.NewReplacer("yyyy", "2006", "yy", "06", "MMM", "Jan", "MM", "01", "dd", "02",
		"HH", "15", "mm", "04", "ss", "05", "SSS", "000", "SS", "00", "Z", "-0700")

	layout = replacer.Replace(layout)

	return layout
}

// ParseTime
// @param  value   string  文本格式时间
// @param  layout  string  时间格式
func ParseTime(value string, layout string) time.Time {
	timeVal, _ := time.ParseInLocation(ConvetTimeLayout(layout), value, time.Local)
	return timeVal
}

// 格式化时间文本
// @param  t  			time.Time  	时间
// @param  layout   string  		时间格式
func FormatTime(t time.Time, layout string) string {
	return t.Format(ConvetTimeLayout(layout))
}

// 计算时间持续时间
func Duration(beginTime time.Time, endTime time.Time) time.Duration {
	return endTime.Sub(beginTime)
}
