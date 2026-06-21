package queues_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/20-queues"
)

func TestNew(t *testing.T) {
	ds := queues.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
