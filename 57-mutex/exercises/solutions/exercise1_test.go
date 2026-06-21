package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/57-mutex/exercises/solutions"
)

func TestSafeCounter(t *testing.T) {
	c := &solutions.SafeCounter{}
	c.Inc()
	c.Inc()
	if c.Value() != 2 {
		t.Fatal("expected 2")
	}
}
