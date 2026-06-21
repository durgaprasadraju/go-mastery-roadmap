package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/04-functions/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.Compose(func(x int) int { return x * 2 }, func(x int) int { return x + 1 })(3); got != 8 {
		t.Fatalf("got %v want %v", got, 8)
	}
	if got := len(solutions.MapInts([]int{1,2}, func(x int) int { return x * 2 })); got != 2 {
		t.Fatalf("got %v want %v", got, 2)
	}
}
