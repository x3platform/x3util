package pathutil

import (
	"os"
	"path/filepath"
)

// Check the path exists
func Exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func GetFileDir(path string) string {
	return filepath.Dir(path)
}
