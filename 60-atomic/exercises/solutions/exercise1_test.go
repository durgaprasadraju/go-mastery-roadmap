package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/60-atomic/exercises/solutions"
)

func TestAtomicCounter(t *testing.T) {
	c := &solutions.AtomicCounter{}
	c.Add(5)
	if c.Load() != 5 {
		t.Fatal("expected 5")
	}
}
