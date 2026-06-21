package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/38-backtracking/exercises/solutions"
)

func TestPermute(t *testing.T) {
	if len(solutions.Permute([]int{1,2,3})) != 6 {
		t.Fatal("expected 6 permutations")
	}
}
