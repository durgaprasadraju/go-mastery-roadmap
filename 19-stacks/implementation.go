// Package stacks implements a stack (LIFO) data structure from scratch.
package stacks

import "errors"

var ErrEmpty = errors.New("stack is empty")

// Stack is a slice-backed LIFO stack. Amortized O(1) push/pop.
type Stack struct {
	items []int
}

// New creates an empty stack.
func New() *Stack {
	return &Stack{items: make([]int, 0)}
}

// Push adds an element. O(1) amortized.
func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

// Pop removes and returns the top element. O(1).
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmpty
	}
	top := len(s.items) - 1
	v := s.items[top]
	s.items = s.items[:top]
	return v, nil
}

// Peek returns the top without removing. O(1).
func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmpty
	}
	return s.items[len(s.items)-1], nil
}

// Size returns element count.
func (s *Stack) Size() int { return len(s.items) }

// IsEmpty reports whether the stack is empty.
func (s *Stack) IsEmpty() bool { return len(s.items) == 0 }
