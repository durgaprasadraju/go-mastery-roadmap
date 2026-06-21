package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/51-worker-pools/exercises/solutions"
)

func TestRunPool(t *testing.T) {
	if solutions.RunPool(10, 4) != 10 {
		t.Fatal("expected 10")
	}
}
