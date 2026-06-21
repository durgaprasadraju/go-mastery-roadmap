package heaps_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/23-heaps"
)

func TestNew(t *testing.T) {
	ds := heaps.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
