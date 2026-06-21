package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/22-binary-search-tree/exercises/solutions"
)

func TestBSTInsertSearch(t *testing.T) {
	root := solutions.InsertBST(nil, 5)
	root = solutions.InsertBST(root, 3)
	root = solutions.InsertBST(root, 7)
	if !solutions.SearchBST(root, 3) || solutions.SearchBST(root, 6) {
		t.Fatal("bst search failed")
	}
}
