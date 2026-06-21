package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/114-event-driven/exercises/solutions"
)

func TestOrderEvents(t *testing.T) {
	o := &solutions.OrderWithEvents{ID: "1"}
	o.Ship()
	if len(o.Events()) != 1 || o.Events()[0].Name != "OrderShipped" {
		t.Fatal("expected OrderShipped event")
	}
}
