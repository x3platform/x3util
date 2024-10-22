package jsonutil

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func Stringify(v any) string {
	// Marshal result is byte[]
	buf, err := json.Marshal(v)
	if err != nil {
		log.Error(err.Error())
	}
	return string(buf)
}
