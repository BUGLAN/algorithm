package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	var ok bool
	circularQueue := Constructor(3)
	ok = circularQueue.EnQueue(1)
	assert.True(t, ok)
	// circularQueue.EnQueue(2)
	// assert.True(t, ok)
	// circularQueue.EnQueue(3)
	// assert.True(t, ok)
	// circularQueue.EnQueue(4)
	// assert.True(t, ok)
	// circularQueue.Rear()
	// circularQueue.IsFull()
	// circularQueue.DeQueue()
	// circularQueue.EnQueue(4)
	// circularQueue.Rear()
}
