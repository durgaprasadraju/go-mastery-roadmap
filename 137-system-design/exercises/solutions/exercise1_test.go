package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/137-system-design/exercises/solutions"
)

func TestURLShortener(t *testing.T) {
	code := solutions.Encode(3)
	if solutions.Decode(code) != 3 {
		t.Fatal("roundtrip failed")
	}
}
