package osutil

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	executablModTime     = time.Now()
	executablModTimeOnce sync.Once
)

// GetExecutableModTime get executable file modified time of current process.
func GetExecutableModTime() time.Time {
	executablModTimeOnce.Do(func() {
		exePath, err := os.Executable()
		if err != nil {
			log.Error("os.Executable: %v", err)
			return
		}

		exePath, err = filepath.Abs(exePath)
		if err != nil {
			log.Error("filepath.Abs: %v", err)
			return
		}

		exePath, err = filepath.EvalSymlinks(exePath)
		if err != nil {
			log.Error("filepath.EvalSymlinks: %v", err)
			return
		}

		st, err := os.Stat(exePath)
		if err != nil {
			log.Error("os.Stat: %v", err)
			return
		}

		executablModTime = st.ModTime()
	})
	return executablModTime
}
