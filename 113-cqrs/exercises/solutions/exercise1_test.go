package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/113-cqrs/exercises/solutions"
)

func TestCQRS(t *testing.T) {
	store := solutions.NewUserStore()
	store.Handle(solutions.CreateUserCommand{ID: "1", Email: "x@y.com"})
	v, ok := store.Get("1")
	if !ok || v.Email != "x@y.com" {
		t.Fatal("projection failed")
	}
}
