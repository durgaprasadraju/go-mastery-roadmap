package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/02-types-system/exercises/solutions"
)

func TestInterviewChallenge(t *testing.T) {
	if got := solutions.InterviewChallengeSolution(); got != 0 {
		t.Fatalf("got %v want 0", got)
	}
}
