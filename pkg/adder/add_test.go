package adder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	const a = 1
	const b = 1
	result := Add(a, b)
	assert.Equal(t, a+b, result)
}
