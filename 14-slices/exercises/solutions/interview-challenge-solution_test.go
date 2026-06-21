package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/14-slices/exercises/solutions"
)

func TestReverseInPlace(t *testing.T) {
	s := []int{1,2,3}
	solutions.ReverseInPlace(s)
	if s[0] != 3 {
		t.Fatalf("got %v", s)
	}
}
