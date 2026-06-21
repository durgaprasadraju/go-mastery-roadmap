package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/14-slices/exercises/solutions"
)

func TestAppendUnique(t *testing.T) {
	got := solutions.AppendUnique([]int{1}, 1)
	if len(got) != 1 {
		t.Fatal("duplicate added")
	}
}
