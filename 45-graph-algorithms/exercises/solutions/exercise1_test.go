package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/45-graph-algorithms/exercises/solutions"
)

func TestBFS(t *testing.T) {
	g := solutions.Graph{1: {2}, 2: {3}}
	order := g.BFS(1)
	if len(order) != 3 {
		t.Fatalf("got %v", order)
	}
}
