package searching_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/32-searching"
)

func TestNew(t *testing.T) {
	ds := searching.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
