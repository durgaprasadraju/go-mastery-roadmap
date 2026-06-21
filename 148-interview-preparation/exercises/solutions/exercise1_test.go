package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/148-interview-preparation/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := (solutions.ModuleProgress{"x",100}).Completed(); got != true {
		t.Fatalf("got %v want %v", got, true)
	}
	if got := (solutions.ModuleProgress{"x",50}).Completed(); got != false {
		t.Fatalf("got %v want %v", got, false)
	}
}
