package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/83-database-sql/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := solutions.BuildSelect("users",[]string{"id","name"},"id=$1"); got != "SELECT id, name FROM users WHERE id=$1" {
		t.Fatalf("got %v want %v", got, "SELECT id, name FROM users WHERE id=$1")
	}
	if got := solutions.QuoteIdentifier("order"); got != "\"order\"" {
		t.Fatalf("got %v want %v", got, "\"order\"")
	}
}
