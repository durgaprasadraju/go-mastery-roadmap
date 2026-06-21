package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/43-math-algorithms/exercises/solutions"
)

func TestGCD(t *testing.T) {
	if solutions.GCD(12, 8) != 4 {
		t.Fatal("expected 4")
	}
}
