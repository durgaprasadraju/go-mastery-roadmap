package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/99-security/exercises/solutions"
)

func TestParseBearer(t *testing.T) {
	tok, ok := solutions.ParseBearer("Bearer abc123")
	if !ok || tok != "abc123" {
		t.Fatalf("tok=%q ok=%v", tok, ok)
	}
}
