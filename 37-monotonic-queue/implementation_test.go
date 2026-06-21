package monotonic_queue_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/37-monotonic-queue"
)

func TestNew(t *testing.T) {
	ds := monotonic_queue.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
