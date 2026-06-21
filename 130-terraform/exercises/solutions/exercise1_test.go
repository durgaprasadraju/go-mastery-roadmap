package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/130-terraform/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.Liveness().OK; got != true {
		t.Fatalf("got %v want %v", got, true)
	}
	if got := solutions.Readiness(true,true).OK; got != true {
		t.Fatalf("got %v want %v", got, true)
	}
}
