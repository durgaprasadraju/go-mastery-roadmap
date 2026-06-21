package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/41-dynamic-programming/exercises/solutions"
)

func TestLIS(t *testing.T) {
	if solutions.LIS([]int{10,9,2,5,3,7,101,18}) != 4 {
		t.Fatal("expected LIS length 4")
	}
}
