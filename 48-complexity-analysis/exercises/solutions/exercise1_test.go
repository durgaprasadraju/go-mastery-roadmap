package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/48-complexity-analysis/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.CountOperations(10); got != 10 {
		t.Fatalf("got %v want %v", got, 10)
	}
	if got := solutions.ClassifyGrowth(100,"linear"); got != "O(n)" {
		t.Fatalf("got %v want %v", got, "O(n)")
	}
}
