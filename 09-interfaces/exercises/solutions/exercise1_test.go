package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/09-interfaces/exercises/solutions"
)

func TestPersonString(t *testing.T) {
	p := solutions.Person{"Ada", "Lovelace"}
	if p.String() != "Lovelace, Ada" {
		t.Fatalf("got %q", p.String())
	}
}
