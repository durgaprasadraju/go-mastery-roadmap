package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/40-greedy/exercises/solutions"
)

func TestMaxActivities(t *testing.T) {
	if solutions.MaxActivities([]int{1,2,3}) < 1 {
		t.Fatal("expected at least 1")
	}
}
