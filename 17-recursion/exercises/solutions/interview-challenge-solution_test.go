package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/17-recursion/exercises/solutions"
)

func TestFib(t *testing.T) {
	if solutions.Fib(10) != 55 {
		t.Fatal("expected 55")
	}
}
