package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/134-nats/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.NewMessage("orders","k",[]byte("x")).Topic; got != "orders" {
		t.Fatalf("got %v want %v", got, "orders")
	}
}
