package solutions

type TreeNode struct{ Val int; Left, Right *TreeNode }

// InsertBST inserts val into BST rooted at node.
func InsertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = InsertBST(root.Left, val)
	} else if val > root.Val {
		root.Right = InsertBST(root.Right, val)
	}
	return root
}

// SearchBST returns true if val exists in BST.
func SearchBST(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.Val {
		return true
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}