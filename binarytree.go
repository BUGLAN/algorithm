package algorithm

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 前序遍历: 先遍历根节点, 再遍历左子树
func PreOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	PreOrder(root.Left, list)
	PreOrder(root.Right, list)
}

// InOrder 中序遍历: 先遍历左子树, 再遍历根节点, 再遍历左子树
func InOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, list)
	*list = append(*list, root.Val)
	InOrder(root.Right, list)
}

// BackOrder 后序遍历: 先遍历左子树, 再遍右子树, 再遍历根
func BackOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	BackOrder(root.Left, list)
	BackOrder(root.Right, list)
	*list = append(*list, root.Val)
}
