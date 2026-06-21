package stacks_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/19-stacks"
)

func TestStackLIFO(t *testing.T) {
	s := stacks.New()
	s.Push(1)
	s.Push(2)
	v, err := s.Pop()
	if err != nil || v != 2 {
		t.Fatalf("pop = %d err = %v", v, err)
	}
	v, err = s.Peek()
	if err != nil || v != 1 {
		t.Fatalf("peek = %d", v)
	}
}

func TestStackEmpty(t *testing.T) {
	s := stacks.New()
	if _, err := s.Pop(); err != stacks.ErrEmpty {
		t.Fatal("expected ErrEmpty")
	}
}

func BenchmarkPushPop(b *testing.B) {
	s := stacks.New()
	for i := 0; i < b.N; i++ {
		s.Push(i)
		s.Pop()
	}
}
