package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/17-recursion/exercises/solutions"
)

func TestFactorial(t *testing.T) {
	if solutions.Factorial(5) != 120 {
		t.Fatal("expected 120")
	}
}
