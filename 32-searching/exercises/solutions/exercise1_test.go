package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/32-searching/exercises/solutions"
)

func TestBinarySearch(t *testing.T) {
	if solutions.BinarySearch([]int{1,3,5}, 3) != 1 {
		t.Fatal("expected index 1")
	}
}
