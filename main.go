package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"x3platform.com/x3util/cmd"
)

// @title util
// @version 0.0.5
// @description x3util
// @license.name MIT
// @license.url https://github.com/x3platform/x3util/blob/main/LICENSE.md
func main() {
	for _, arg := range os.Args {
		if arg == "--verbose" {
			log.SetLevel(log.TraceLevel)
		}
	}

	// 设置输出日志格式
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.Debugf("version:%s", "1.0.0-build")
	log.Debugf("commit-id:%s", "git-commit-id-build")
	log.Debugf("build date:%s", "make-date-build")
	log.Debugf("log level:%s", log.GetLevel())

	cmd.Execute()
}
