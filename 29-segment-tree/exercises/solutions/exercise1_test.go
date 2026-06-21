package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/29-segment-tree/exercises/solutions"
)

func TestRangeSum(t *testing.T) {
	if solutions.RangeSum([]int{1,2,3}) != 6 {
		t.Fatal("expected 6")
	}
}
