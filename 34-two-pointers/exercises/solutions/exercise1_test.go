package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/34-two-pointers/exercises/solutions"
)

func TestTwoSumSorted(t *testing.T) {
	p := solutions.TwoSumSorted([]int{1,2,3,4}, 7)
	if p != [2]int{2,3} {
		t.Fatalf("got %v", p)
	}
}
