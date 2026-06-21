package solutions

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Prepend adds value at head.
func Prepend(head *ListNode, val int) *ListNode {
	return &ListNode{Val: val, Next: head}
}