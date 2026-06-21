package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/06-pointers/exercises/solutions"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	solutions.Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Fatalf("a=%d b=%d", a, b)
	}
}
