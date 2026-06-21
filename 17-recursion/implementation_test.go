package recursion_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/17-recursion"
)

func TestNew(t *testing.T) {
	ds := recursion.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
