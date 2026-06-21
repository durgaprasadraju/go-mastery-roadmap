package sliding_window_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/33-sliding-window"
)

func TestNew(t *testing.T) {
	ds := sliding_window.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
