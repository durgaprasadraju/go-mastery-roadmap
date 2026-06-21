package solutions

import "errors"

var ErrEmpty = errors.New("stack empty")

// Stack is a LIFO stack.
type Stack struct{ items []int }

func NewStack() *Stack { return &Stack{} }

func (s *Stack) Push(v int) { s.items = append(s.items, v) }

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmpty
	}
	i := len(s.items) - 1
	v := s.items[i]
	s.items = s.items[:i]
	return v, nil
}