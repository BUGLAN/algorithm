package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	list := make([]int, 0)
	PreOrder(root, &list)
	assert.Equal(t, 3, len(list))
	assert.Equal(t, []int{1, 2, 3}, list)
}

func TestInOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	list := make([]int, 0)
	InOrder(root, &list)
	assert.Equal(t, 3, len(list))
	assert.Equal(t, []int{2, 1, 3}, list)
}

func TestBackOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	list := make([]int, 0)
	BackOrder(root, &list)
	assert.Equal(t, 3, len(list))
	assert.Equal(t, []int{2, 3, 1}, list)
}
