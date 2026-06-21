package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/10-generics/exercises/solutions"
)

func TestMax(t *testing.T) {
	if solutions.MaxInt(3, 7) != 7 {
		t.Fatal("expected 7")
	}
}
