package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/51-worker-pools/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	got := solutions.InterviewChallengeSolution([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if got != 6 {
		t.Fatalf("got %d want 6", got)
	}
}
