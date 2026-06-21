package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/19-stacks/exercises/solutions"
)

func TestValidParentheses(t *testing.T) {
	if !solutions.ValidParentheses("()[]{}") || solutions.ValidParentheses("(]") {
		t.Fatal("parens check failed")
	}
}
