package segment_tree_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/29-segment-tree"
)

func TestNew(t *testing.T) {
	ds := segment_tree.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
