package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/07-structs/exercises/solutions"
)

func TestRectangle(t *testing.T) {
	r := solutions.Rectangle{3, 4}
	if r.Area() != 12 || r.Perimeter() != 14 {
		t.Fatal("wrong area or perimeter")
	}
}
