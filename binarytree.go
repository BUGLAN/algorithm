package algorithm

// TreeNode 二叉树节点
type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode create tree node
func NewTreeNode(val interface{}) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// PreOrder 前序遍历
func (root *TreeNode) PreOrder() {

}

// InOrder 中序遍历
func (root *TreeNode) InOrder() {

}

// BackOrder 后序遍历
func (root *TreeNode) BackOrder() {

}
