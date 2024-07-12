package arrayutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	var result bool

	arr := []string{"aa", "bb", "cc", "dd"}

	result = Contains(arr, "aa")
	assert.Equal(t, true, result)

	result = Contains(arr, "a")
	assert.Equal(t, false, result)

	result = Contains(arr, "e")
	assert.Equal(t, false, result)
}
