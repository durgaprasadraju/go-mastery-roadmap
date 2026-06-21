package bit_manipulation_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/42-bit-manipulation"
)

func TestNew(t *testing.T) {
	ds := bit_manipulation.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
