package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/112-ddd/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution(3, 2) {
		t.Fatal("expected valid order")
	}
	if solutions.InterviewChallengeSolution(1, 2) {
		t.Fatal("expected invalid order")
	}
}
