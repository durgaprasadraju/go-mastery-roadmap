package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/95-authorization/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := len(solutions.HashPasswordSHA256("secret")); got != 64 {
		t.Fatalf("got %v want %v", got, 64)
	}
	if got := solutions.ConstantTimeCompare("abc","abc"); got != true {
		t.Fatalf("got %v want %v", got, true)
	}
}
