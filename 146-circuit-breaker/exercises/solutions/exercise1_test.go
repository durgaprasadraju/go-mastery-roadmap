package solutions_test

import (
	"errors"
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/146-circuit-breaker/exercises/solutions"
)

func TestCircuitBreaker(t *testing.T) {
	cb := solutions.NewCircuitBreaker(2)
	errFn := func() error { return errors.New("fail") }
	_ = cb.Call(errFn)
	if err := cb.Call(errFn); err == nil {
		t.Fatal("expected open circuit")
	}
}
