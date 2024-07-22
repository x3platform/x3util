package pathutil

import (
	"os"
)

// 路径是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// func GetFiles(folder string) {
// 	files, _ := ioutil.ReadDir(folder)
// 	for _, file := range files {
// 		if file.IsDir() {
// 			GetFiles(folder + "/" + file.Name())
// 		} else {
// 			fmt.Println(folder + "/" + file.Name())
// 		}
// 	}
// }
