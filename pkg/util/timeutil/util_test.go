package timeutil

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

const TEST_TIME_TEXT = "2020-12-12 12:12:20"
const TEST_TIME_LOCAL_LAYOUT = "2006-01-02 15:04:05"
const TEST_TIME_UNIX_TIMESTAMP = 1607746340

func TestTime(t *testing.T) {
	var timeValue time.Time = Time(TEST_TIME_UNIX_TIMESTAMP)

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	timeTemplate2 := "2006/01/02 15:04:05" //其他类型
	timeTemplate3 := "2006-01-02"          //其他类型
	timeTemplate4 := "15:04:05"            //其他类型

	// ======= 将时间戳格式化为日期字符串 =======
	log.Println(time.Unix(TEST_TIME_UNIX_TIMESTAMP, 0).Format(timeTemplate1)) //输出：2019-01-08 13:50:30
	log.Println(time.Unix(TEST_TIME_UNIX_TIMESTAMP, 0).Format(timeTemplate2)) //输出：2019/01/08 13:50:30
	log.Println(time.Unix(TEST_TIME_UNIX_TIMESTAMP, 0).Format(timeTemplate3)) //输出：2019-01-08
	log.Println(time.Unix(TEST_TIME_UNIX_TIMESTAMP, 0).Format(timeTemplate4)) //输出：13:50:30

	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, TEST_TIME_TEXT, time.Local) // 使用 parseInLocation 将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())                                                   // 输出：1546926630

	assert.Equal(t, int64(TEST_TIME_UNIX_TIMESTAMP), timeValue.Unix())
}

func TestParseTime(t *testing.T) {
	var (
		timeValue time.Time
		timeText  string
	)
	timeValue = ParseTime(TEST_TIME_TEXT, TEST_TIME_LOCAL_LAYOUT)
	// RFC3339 = "2006-01-02T15:04:05Z07:00"
	timeText = timeValue.Format(time.RFC3339)
	assert.Equal(t, "2020-12-12T12:12:20+08:00", timeText)
}

func TestFormatTime(t *testing.T) {
	var (
		timeValue time.Time
		timeText  string
	)

	timeValue = ParseTime(TEST_TIME_TEXT, TEST_TIME_LOCAL_LAYOUT)

	timeText = FormatTime(timeValue, "yyyy-MM-ddTHH:mm:ssZ")
	assert.Equal(t, "2020-12-12T12:12:20+0800", timeText)

	timeText = FormatTime(timeValue, "yyyy-MM-dd HH:mm:ss")
	assert.Equal(t, "2020-12-12 12:12:20", timeText)

	timeText = FormatTime(timeValue, "yyyy-MM-dd")
	assert.Equal(t, "2020-12-12", timeText)

	timeText = FormatTime(timeValue, "HH:mm:ss")
	assert.Equal(t, "12:12:20", timeText)
}

func TestDuration(t *testing.T) {
	var (
		timeValue time.Time
		timeText  string
	)

	timeValue = ParseTime(TEST_TIME_TEXT, TEST_TIME_LOCAL_LAYOUT)

	timeText = FormatTime(timeValue, "yyyy-MM-ddTHH:mm:ssZ")
	assert.Equal(t, "2020-12-12T12:12:20+0800", timeText)

	timeText = FormatTime(timeValue, "yyyy-MM-dd HH:mm:ss")
	assert.Equal(t, "2020-12-12 12:12:20", timeText)

	timeText = FormatTime(timeValue, "yyyy-MM-dd")
	assert.Equal(t, "2020-12-12", timeText)

	timeText = FormatTime(timeValue, "HH:mm:ss")
	assert.Equal(t, "12:12:20", timeText)
}
