package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/03-control-flow/exercises/solutions"
)

func TestFizzBuzz(t *testing.T) {
	got := solutions.FizzBuzz(5)
	want := []string{"1","2","Fizz","4","Buzz"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at %d: got %q want %q", i, got[i], want[i])
		}
	}
}
