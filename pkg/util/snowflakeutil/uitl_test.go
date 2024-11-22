package snowflakeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewId(t *testing.T) {
	id := NewId()

	assert.True(t, id > 100000000000000000)
}

func TestNewIdString(t *testing.T) {
	text := NewIdString()

	assert.Equal(t, 18, len(text))
}
