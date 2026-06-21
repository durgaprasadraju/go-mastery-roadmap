package maps_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/15-maps"
)

func TestNew(t *testing.T) {
	ds := maps.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
