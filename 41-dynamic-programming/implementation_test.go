package dynamic_programming_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/41-dynamic-programming"
)

func TestNew(t *testing.T) {
	ds := dynamic_programming.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
