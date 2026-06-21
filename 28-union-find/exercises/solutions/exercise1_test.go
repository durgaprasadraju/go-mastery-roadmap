package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/28-union-find/exercises/solutions"
)

func TestUnionFind(t *testing.T) {
	uf := solutions.NewUF(3)
	uf.Union(0,1)
	if uf.Find(0) != uf.Find(1) {
		t.Fatal("expected connected")
	}
}
