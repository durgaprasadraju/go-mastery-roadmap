package union_find_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/28-union-find"
)

func TestNew(t *testing.T) {
	ds := union_find.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
