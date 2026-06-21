package arrays_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/13-arrays"
)

func TestNew(t *testing.T) {
	ds := arrays.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
