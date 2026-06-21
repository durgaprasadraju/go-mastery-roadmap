package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/112-ddd/exercises/solutions"
)

func TestOrderAggregate(t *testing.T) {
	o, err := solutions.NewOrder("o1", []string{"item"})
	if err != nil || o.Status != "pending" {
		t.Fatalf("order=%+v err=%v", o, err)
	}
	if err := o.Ship(); err != nil || o.Status != "shipped" {
		t.Fatalf("ship failed: %v", err)
	}
}
