package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/103-testing/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.Add(2,3); got != 5 {
		t.Fatalf("got %v want %v", got, 5)
	}
	if got := solutions.IsEven(4); got != true {
		t.Fatalf("got %v want %v", got, true)
	}
}
