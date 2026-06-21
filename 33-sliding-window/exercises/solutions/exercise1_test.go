package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/33-sliding-window/exercises/solutions"
)

func TestMaxSumSubarrayK(t *testing.T) {
	if solutions.MaxSumSubarrayK([]int{1,2,3,4,5}, 2) != 9 {
		t.Fatal("expected 9")
	}
}
