package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/16-strings-runes-bytes/exercises/solutions"
)

func TestIsPalindrome(t *testing.T) {
	if !solutions.IsPalindrome("Level") {
		t.Fatal("expected palindrome")
	}
}
