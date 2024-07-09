package boolutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	var val bool

	val = Bool("1")
	assert.Equal(t, true, val)

	val = Bool("0")
	assert.Equal(t, false, val)

	val = Bool("True")
	assert.Equal(t, true, val)

	val = Bool("False")
	assert.Equal(t, false, val)

	val = Bool("Yes")
	assert.Equal(t, true, val)

	val = Bool("YES")
	assert.Equal(t, true, val)

	val = Bool("No")
	assert.Equal(t, false, val)

	val = Bool("On")
	assert.Equal(t, true, val)

	val = Bool("ON")
	assert.Equal(t, true, val)

	val = Bool("Off")
	assert.Equal(t, false, val)
}
