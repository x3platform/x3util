package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUpper(t *testing.T) {
	text := FirstUpper("abc")

	assert.Equal(t, "Abc", text)
}

func TestFirstLower(t *testing.T) {
	text := FirstLower("ABC")

	assert.Equal(t, "aBC", text)
}

func TestRandom(t *testing.T) {
	text := Random(20)

	assert.Len(t, text, 20)
}
