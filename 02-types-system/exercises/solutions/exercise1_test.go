package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/02-types-system/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.TypeName(42); got != "int" {
		t.Fatalf("got %v want %v", got, "int")
	}
	if got := solutions.IsNumeric(42); got != true {
		t.Fatalf("got %v want %v", got, true)
	}
}
