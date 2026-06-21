package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/31-sorting/exercises/solutions"
)

func TestSortColors(t *testing.T) {
	a := []int{2,0,2,1,1,0}
	solutions.SortColors(a)
	if a[0] != 0 || a[len(a)-1] != 2 {
		t.Fatalf("got %v", a)
	}
}
