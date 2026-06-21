package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/55-select/exercises/solutions"
)

func TestParallelSum(t *testing.T) {
	if solutions.ParallelSum([]int{1,2,3,4}, 2) != 10 {
		t.Fatal("expected 10")
	}
}
