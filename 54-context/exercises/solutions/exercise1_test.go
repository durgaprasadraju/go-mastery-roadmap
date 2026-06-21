package solutions_test

import (
	"context"
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/54-context/exercises/solutions"
)

func TestWithTimeoutOp(t *testing.T) {
	if err := solutions.WithTimeoutOp(context.Background(), 1); err == nil {
		t.Fatal("expected timeout")
	}
}
