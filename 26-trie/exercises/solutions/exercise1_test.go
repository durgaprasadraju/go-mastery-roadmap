package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/26-trie/exercises/solutions"
)

func TestTrie(t *testing.T) {
	tr := solutions.NewTrie()
	tr.Insert("go")
	if !tr.Search("go") {
		t.Fatal("expected found")
	}
}
