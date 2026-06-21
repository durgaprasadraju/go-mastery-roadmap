package priority_queue_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/24-priority-queue"
)

func TestNew(t *testing.T) {
	ds := priority_queue.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
