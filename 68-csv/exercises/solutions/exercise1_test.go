package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/68-csv/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := len(solutions.Lines("a\n\nb")); got != 2 {
		t.Fatalf("got %v want %v", got, 2)
	}
	if got := len(solutions.Lines("")); got != 0 {
		t.Fatalf("got %v want %v", got, 0)
	}
}
