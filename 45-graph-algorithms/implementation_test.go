package graph_algorithms_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/45-graph-algorithms"
)

func TestNew(t *testing.T) {
	ds := graph_algorithms.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
