package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/145-rate-limiting/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	sw := solutions.NewSlidingWindow(2)
	if !sw.Allow() || !sw.Allow() || sw.Allow() {
		t.Fatal("sliding window failed")
	}
}
