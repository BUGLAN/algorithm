package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = Stack{}

func init() {
	s.Push("1")
	s.Push(1)
	s.Push(2)
}

func TestStack(t *testing.T) {
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, 2, s.Pop().(int))
}
