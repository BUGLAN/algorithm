package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &Stack{}

func init() {
	s.Push("1")
	s.Push(1)
	s.Push(2)
}

func TestStack(t *testing.T) {
	assert.Equal(t, 3, s.Len())
	value, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 2, value)
}

func TestStack_Empty(t *testing.T) {
	s = &Stack{}
	value, err := s.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, value)

}
