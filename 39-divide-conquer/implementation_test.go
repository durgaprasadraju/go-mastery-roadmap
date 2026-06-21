package divide_conquer_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/39-divide-conquer"
)

func TestNew(t *testing.T) {
	ds := divide_conquer.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
