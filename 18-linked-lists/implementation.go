// Package linkedlists implements singly, doubly, and circular linked lists from scratch.
package linkedlists

import "errors"

var (
	ErrEmptyList = errors.New("linked list is empty")
	ErrNotFound  = errors.New("element not found")
)

// Node represents a singly linked list node.
type Node struct {
	Value int
	Next  *Node
}

// SinglyLinkedList is a singly linked list.
type SinglyLinkedList struct {
	head *Node
	size int
}

// NewSingly creates an empty singly linked list.
func NewSingly() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Size returns the number of elements. O(1) with cached counter.
func (l *SinglyLinkedList) Size() int { return l.size }

// IsEmpty reports whether the list has no elements.
func (l *SinglyLinkedList) IsEmpty() bool { return l.size == 0 }

// Prepend inserts at the head. O(1).
func (l *SinglyLinkedList) Prepend(v int) {
	l.head = &Node{Value: v, Next: l.head}
	l.size++
}

// Append inserts at the tail. O(n).
func (l *SinglyLinkedList) Append(v int) {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
		l.size++
		return
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
	l.size++
}

// Remove removes the first occurrence of v. O(n).
func (l *SinglyLinkedList) Remove(v int) error {
	if l.head == nil {
		return ErrEmptyList
	}
	if l.head.Value == v {
		l.head = l.head.Next
		l.size--
		return nil
	}
	cur := l.head
	for cur.Next != nil {
		if cur.Next.Value == v {
			cur.Next = cur.Next.Next
			l.size--
			return nil
		}
		cur = cur.Next
	}
	return ErrNotFound
}

// ToSlice returns all values in order. O(n).
func (l *SinglyLinkedList) ToSlice() []int {
	out := make([]int, 0, l.size)
	for cur := l.head; cur != nil; cur = cur.Next {
		out = append(out, cur.Value)
	}
	return out
}

// DoublyNode represents a doubly linked list node.
type DoublyNode struct {
	Value int
	Prev  *DoublyNode
	Next  *DoublyNode
}

// DoublyLinkedList is a doubly linked list with O(1) append/prepend.
type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
	size int
}

// NewDoubly creates an empty doubly linked list.
func NewDoubly() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Size returns element count.
func (l *DoublyLinkedList) Size() int { return l.size }

// Append adds to tail. O(1).
func (l *DoublyLinkedList) Append(v int) {
	n := &DoublyNode{Value: v}
	if l.tail == nil {
		l.head, l.tail = n, n
	} else {
		n.Prev = l.tail
		l.tail.Next = n
		l.tail = n
	}
	l.size++
}

// Prepend adds to head. O(1).
func (l *DoublyLinkedList) Prepend(v int) {
	n := &DoublyNode{Value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.Next = l.head
		l.head.Prev = n
		l.head = n
	}
	l.size++
}

// Remove removes first occurrence. O(n).
func (l *DoublyLinkedList) Remove(v int) error {
	for cur := l.head; cur != nil; cur = cur.Next {
		if cur.Value == v {
			if cur.Prev != nil {
				cur.Prev.Next = cur.Next
			} else {
				l.head = cur.Next
			}
			if cur.Next != nil {
				cur.Next.Prev = cur.Prev
			} else {
				l.tail = cur.Prev
			}
			l.size--
			return nil
		}
	}
	if l.size == 0 {
		return ErrEmptyList
	}
	return ErrNotFound
}

// CircularLinkedList is a circular singly linked list.
type CircularLinkedList struct {
	head *Node
	size int
}

// NewCircular creates an empty circular list.
func NewCircular() *CircularLinkedList {
	return &CircularLinkedList{}
}

// Append inserts at end of circle. O(n) without tail pointer.
func (l *CircularLinkedList) Append(v int) {
	n := &Node{Value: v}
	if l.head == nil {
		n.Next = n
		l.head = n
	} else {
		tail := l.head
		for tail.Next != l.head {
			tail = tail.Next
		}
		tail.Next = n
		n.Next = l.head
	}
	l.size++
}

// Traverse visits each node once. O(n).
func (l *CircularLinkedList) Traverse() []int {
	if l.head == nil {
		return nil
	}
	out := make([]int, 0, l.size)
	cur := l.head
	for {
		out = append(out, cur.Value)
		cur = cur.Next
		if cur == l.head {
			break
		}
	}
	return out
}

// Size returns element count.
func (l *CircularLinkedList) Size() int { return l.size }
