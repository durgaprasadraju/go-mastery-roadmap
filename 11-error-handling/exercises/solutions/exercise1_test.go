package solutions_test

import (
	"errors"
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/11-error-handling/exercises/solutions"
)

func TestDivide(t *testing.T) {
	v, err := solutions.Divide(10, 2)
	if err != nil || v != 5 {
		t.Fatalf("v=%d err=%v", v, err)
	}
	_, err = solutions.Divide(1, 0)
	if !errors.Is(err, solutions.ErrDivZero) {
		t.Fatal("expected ErrDivZero")
	}
}
