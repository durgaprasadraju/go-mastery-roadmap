package solutions

// IsValidBST checks whether tree satisfies BST property.
func IsValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return IsValidBST(root.Left, min, root.Val) && IsValidBST(root.Right, root.Val, max)
}