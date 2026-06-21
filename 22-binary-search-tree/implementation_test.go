package binary_search_tree_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/22-binary-search-tree"
)

func TestNew(t *testing.T) {
	ds := binary_search_tree.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
