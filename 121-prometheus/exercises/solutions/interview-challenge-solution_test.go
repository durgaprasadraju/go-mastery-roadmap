package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/121-prometheus/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution(4) {
		t.Fatal("expected even")
	}
	if solutions.InterviewChallengeSolution(3) {
		t.Fatal("expected odd")
	}
}
