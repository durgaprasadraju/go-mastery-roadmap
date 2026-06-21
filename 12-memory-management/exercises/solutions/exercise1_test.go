package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/12-memory-management/exercises/solutions"
)

func TestSharesBacking(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[:2]
	if !solutions.SharesBacking(a, b) {
		t.Fatal("expected shared backing")
	}
}
