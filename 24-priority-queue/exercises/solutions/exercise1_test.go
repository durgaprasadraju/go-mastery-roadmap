package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/24-priority-queue/exercises/solutions"
)

func TestStack(t *testing.T) {
	s := solutions.NewStack()
	s.Push(1)
	s.Push(2)
	v, err := s.Pop()
	if err != nil || v != 2 {
		t.Fatalf("v=%d err=%v", v, err)
	}
}
