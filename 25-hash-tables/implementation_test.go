package hash_tables_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/25-hash-tables"
)

func TestNew(t *testing.T) {
	ds := hash_tables.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
