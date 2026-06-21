package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/15-maps/exercises/solutions"
)

func TestWordCount(t *testing.T) {
	m := solutions.WordCount([]string{"a","b","a"})
	if m["a"] != 2 {
		t.Fatal("expected count 2")
	}
}
