package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/71-networking-fundamentals/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution("GET", "/api/v1/users", "GET", "/api") {
		t.Fatal("route should match")
	}
}
