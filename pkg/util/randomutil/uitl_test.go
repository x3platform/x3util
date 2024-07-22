package randomutil

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"x3platform.com/x3util/pkg/util/arrayutil"
)

func TestUUID(t *testing.T) {
	var keys = []string{}
	for i := 0; i < 100; i++ {
		text := UUID()
		log.Info(text)
		assert.Len(t, text, 36)
		if arrayutil.Contains(keys, text) {
			assert.Fail(t, "fail", text)
		}
		keys = append(keys, text)
	}
}

func TestUUID32(t *testing.T) {
	var keys = []string{}
	for i := 0; i < 100; i++ {
		text := UUID32()
		log.Info(text)
		assert.Len(t, text, 32)
		if arrayutil.Contains(keys, text) {
			assert.Fail(t, "fail", text)
		}
		keys = append(keys, text)
	}
}

func TestCharacters(t *testing.T) {
	var keys = []string{}
	for i := 0; i < 100; i++ {
		text := Characters(20)
		log.Info(text)
		assert.Len(t, text, 20)
		if arrayutil.Contains(keys, text) {
			assert.Fail(t, "fail", text)
		}
		keys = append(keys, text)
	}
}
