package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/21-trees/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := &solutions.TreeNode{Val: 1, Left: &solutions.TreeNode{Val: 2}}; solutions.TreeHeight(root)
	}
}
