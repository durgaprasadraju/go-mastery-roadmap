package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/18-linked-lists/exercises/solutions"
)

func TestPrepend(t *testing.T) {
	head := solutions.Prepend(nil, 2)
	head = solutions.Prepend(head, 1)
	if head.Val != 1 || head.Next.Val != 2 {
		t.Fatal("prepend failed")
	}
}
