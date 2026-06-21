package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/22-binary-search-tree/exercises/solutions"
)

func TestIsValidBST(t *testing.T) {
	valid := &solutions.TreeNode{
		Val: 2,
		Left: &solutions.TreeNode{Val: 1},
		Right: &solutions.TreeNode{Val: 3},
	}
	if !solutions.IsValidBST(valid, -1<<62, 1<<62) {
		t.Fatal("expected valid BST")
	}
	invalid := &solutions.TreeNode{
		Val: 2,
		Left: &solutions.TreeNode{Val: 3},
	}
	if solutions.IsValidBST(invalid, -1<<62, 1<<62) {
		t.Fatal("expected invalid BST")
	}
}
