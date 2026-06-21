package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/42-bit-manipulation/exercises/solutions"
)

func TestSingleNumber(t *testing.T) {
	if solutions.SingleNumber([]int{2,1,2}) != 1 {
		t.Fatal("expected 1")
	}
}
