package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/15-maps/exercises/solutions"
)

func TestTwoSum(t *testing.T) {
	p := solutions.TwoSum([]int{2,7,11,15}, 9)
	if p != [2]int{0,1} {
		t.Fatalf("got %v", p)
	}
}
