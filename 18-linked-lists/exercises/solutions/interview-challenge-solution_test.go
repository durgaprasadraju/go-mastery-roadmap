package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/18-linked-lists/exercises/solutions"
)

func TestReverseList(t *testing.T) {
	n := solutions.Prepend(solutions.Prepend(nil, 2), 1)
	r := solutions.ReverseList(n)
	if r.Val != 2 {
		t.Fatal("reverse failed")
	}
}
