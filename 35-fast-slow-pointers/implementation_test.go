package fast_slow_pointers_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/35-fast-slow-pointers"
)

func TestNew(t *testing.T) {
	ds := fast_slow_pointers.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
