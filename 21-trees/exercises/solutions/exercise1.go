package solutions

type TreeNode struct{ Val int; Left, Right *TreeNode }

func TreeHeight(root *TreeNode) int {
	if root == nil { return 0 }
	l, r := TreeHeight(root.Left), TreeHeight(root.Right)
	if l > r { return l+1 }
	return r+1
}