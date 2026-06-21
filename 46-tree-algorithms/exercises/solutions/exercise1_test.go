package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/46-tree-algorithms/exercises/solutions"
)

func TestTreeHeight(t *testing.T) {
	root := &solutions.TreeNode{Val:1, Left:&solutions.TreeNode{Val:2}}
	if solutions.TreeHeight(root) != 2 {
		t.Fatal("expected height 2")
	}
}
