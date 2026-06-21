package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/82-load-balancing/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution("GET", "/api/v1/users", "GET", "/api") {
		t.Fatal("route should match")
	}
}
