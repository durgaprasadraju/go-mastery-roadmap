package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/13-arrays/exercises/solutions"
)

func TestFindMax(t *testing.T) {
	if solutions.FindMax([]int{1,5,3}) != 5 {
		t.Fatal("expected 5")
	}
}
