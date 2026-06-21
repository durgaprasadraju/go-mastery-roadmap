package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/58-rwmutex/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if solutions.InterviewChallengeSolution(3) != 3 {
		t.Fatal("expected 3")
	}
}
