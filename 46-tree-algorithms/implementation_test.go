package tree_algorithms_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/46-tree-algorithms"
)

func TestNew(t *testing.T) {
	ds := tree_algorithms.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
