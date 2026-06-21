package exercises_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/01-fundamentals/exercises/solutions"
)

func TestExercisePlaceholder(t *testing.T) {
	// Implement your solution in exercises/ before checking solutions/
	sum, err := solutions.SumInts([]int{1, 2, 3})
	if err != nil || sum != 6 {
		t.Fatalf("verify solutions compile: sum=%d err=%v", sum, err)
	}
}
