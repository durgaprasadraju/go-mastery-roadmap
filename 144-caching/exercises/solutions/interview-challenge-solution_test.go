package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/144-caching/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	n := solutions.InterviewChallengeSolution([]string{"a","b","c"}, 2)
	if n != 2 {
		t.Fatalf("size=%d", n)
	}
}
