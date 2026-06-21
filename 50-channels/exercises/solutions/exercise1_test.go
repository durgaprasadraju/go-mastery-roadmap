package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/50-channels/exercises/solutions"
)

func TestPingPong(t *testing.T) {
	if solutions.PingPong(3) != 3 {
		t.Fatal("expected 3")
	}
}
