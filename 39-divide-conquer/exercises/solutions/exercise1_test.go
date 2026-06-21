package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/39-divide-conquer/exercises/solutions"
)

func TestMergeSort(t *testing.T) {
	a := []int{5,2,8,1}
	solutions.MergeSort(a)
	if a[0] != 1 || a[len(a)-1] != 8 {
		t.Fatalf("got %v", a)
	}
}
