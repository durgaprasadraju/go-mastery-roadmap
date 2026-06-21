package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/41-dynamic-programming/exercises/solutions"
)

func TestCoinChange(t *testing.T) {
	if solutions.CoinChange([]int{1,2,5}, 11) != 3 {
		t.Fatal("expected 3 coins")
	}
}
