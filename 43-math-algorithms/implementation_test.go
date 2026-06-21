package math_algorithms_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/43-math-algorithms"
)

func TestNew(t *testing.T) {
	ds := math_algorithms.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
