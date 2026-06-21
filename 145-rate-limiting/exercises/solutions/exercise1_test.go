package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/145-rate-limiting/exercises/solutions"
)

func TestTokenBucket(t *testing.T) {
	tb := solutions.NewTokenBucket(2)
	if !tb.Allow() || !tb.Allow() || tb.Allow() {
		t.Fatal("token bucket logic failed")
	}
}
