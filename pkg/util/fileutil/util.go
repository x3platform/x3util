package fileutil

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

// Get file name
func GetFileName(path string) string {
	return filepath.Dir(path)
}

// func GetFiles(path string) {
// 	files, _ := ioutil.ReadDir(path)
// 	for _, file := range files {
// 		if file.IsDir() {
// 			GetFiles(path + "/" + file.Name())
// 		} else {
// 			fmt.Println(path + "/" + file.Name())
// 		}
// 	}
// }
