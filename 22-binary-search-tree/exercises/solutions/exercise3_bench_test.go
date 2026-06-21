package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/22-binary-search-tree/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := solutions.InsertBST(nil, 5); solutions.SearchBST(root, 5)
	}
}
