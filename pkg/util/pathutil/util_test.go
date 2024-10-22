package pathutil

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	result, _ := Exists("abc.txt")

	assert.False(t, result)
}

func TestGetFileDir(t *testing.T) {
	result := GetFileDir("/opt/x3util/abc.txt")

	if runtime.GOOS == "windows" {
		assert.Equal(t, "\\opt\\x3util", result, "dir should equal to /opt/x3util")
	} else {
		assert.Equal(t, "/opt/x3util", result, "dir should equal to /opt/x3util")
	}
}
