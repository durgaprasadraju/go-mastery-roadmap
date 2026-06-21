package monotonic_stack_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/36-monotonic-stack"
)

func TestNew(t *testing.T) {
	ds := monotonic_stack.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
