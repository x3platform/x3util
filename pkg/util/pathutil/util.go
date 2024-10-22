package pathutil

import (
	"os"
	"path/filepath"
)

// IsDir returns true if given path is a directory,
// or returns false when it's a file or does not exist.
func IsDir(dir string) (bool, error) {
	f, err := os.Stat(dir)
	if err == nil {
		return f.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) (bool, error) {
	f, err := os.Stat(filePath)
	if err == nil {
		return !f.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Check the path exists
// It returns false when the file or directory does not exist.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Get file directory
func GetFileDir(path string) string {
	return filepath.Dir(path)
}
