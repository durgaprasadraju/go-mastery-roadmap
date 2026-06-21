package trie_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/26-trie"
)

func TestNew(t *testing.T) {
	ds := trie.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
