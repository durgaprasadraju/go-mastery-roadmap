package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/16-strings-runes-bytes/exercises/solutions"
)

func TestRuneLen(t *testing.T) {
	if solutions.RuneLen("café") != 4 {
		t.Fatal("expected 4 runes")
	}
}
