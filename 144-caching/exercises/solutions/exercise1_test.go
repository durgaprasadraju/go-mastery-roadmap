package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/144-caching/exercises/solutions"
)

func TestLRU(t *testing.T) {
	c := solutions.NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatal("expected a=1")
	}
}
