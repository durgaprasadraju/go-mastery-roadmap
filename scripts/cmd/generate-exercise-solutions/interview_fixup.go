package main

import (
	"fmt"
	"strings"
)

func finalizeSpec(m module, spec topicSpec) topicSpec {
	if sameSolutionBody(spec.exercise1, spec.interviewSolution) {
		spec.interviewSolution = genericInterviewSolution(m)
		spec.interviewTest = genericInterviewTest(m)
	} else if sameSolutionBody(spec.exercise1Test, spec.interviewTest) {
		spec.interviewTest = genericInterviewTest(m)
	}
	spec.benchCall = fixBenchCall(m, spec)
	return spec
}

func sameSolutionBody(a, b string) bool {
	return strings.TrimSpace(a) == strings.TrimSpace(b)
}

func genericInterviewSolution(m module) string {
	return `package solutions

// InterviewChallengeSolution returns true when n is even.
func InterviewChallengeSolution(n int) bool {
	return n%2 == 0
}`
}

func genericInterviewTest(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution(4) {
		t.Fatal("expected even")
	}
	if solutions.InterviewChallengeSolution(3) {
		t.Fatal("expected odd")
	}
}
`, m.slug)
}

func fixBenchCall(m module, spec topicSpec) string {
	if b, ok := benchBySlug[m.slug]; ok {
		return b
	}
	bench := strings.TrimSpace(spec.benchCall)
	if strings.Contains(bench, "NewStack") && !strings.Contains(spec.exercise1, "NewStack") {
		if b := benchFromExercise(m.slug, spec.exercise1); b != `_ = 0` {
			return b
		}
	}
	if strings.Contains(bench, "NewOrder") && !strings.Contains(spec.exercise1, "NewOrder") {
		if b := benchFromExercise(m.slug, spec.exercise1); b != `_ = 0` {
			return b
		}
	}
	return prefixBenchStatements(bench)
}

func prefixBenchStatements(bench string) string {
	parts := strings.Split(bench, ";")
	for i, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.Contains(p, ":=") {
			if eq := strings.Index(p, ":="); eq > 0 {
				lhs := strings.TrimSpace(p[:eq])
				rhs := strings.TrimSpace(p[eq+2:])
				if !strings.Contains(rhs, "solutions.") {
					rhs = strings.Replace(rhs, "NewStack()", "solutions.NewStack()", 1)
					rhs = strings.Replace(rhs, "NewQueue()", "solutions.NewQueue()", 1)
					rhs = strings.Replace(rhs, "BankAccount{}", "solutions.BankAccount{}", 1)
					rhs = strings.Replace(rhs, "Rectangle{", "solutions.Rectangle{", 1)
					rhs = strings.Replace(rhs, "TreeNode{", "solutions.TreeNode{", 1)
					rhs = strings.Replace(rhs, "Graph{", "solutions.Graph{", 1)
					rhs = strings.Replace(rhs, "InsertBST(", "solutions.InsertBST(", 1)
					rhs = strings.Replace(rhs, "SearchBST(", "solutions.SearchBST(", 1)
					rhs = strings.Replace(rhs, "TreeHeight(", "solutions.TreeHeight(", 1)
				}
				parts[i] = lhs + " := " + rhs
			}
		} else if strings.Contains(p, ".") && p[0] >= 'a' && p[0] <= 'z' {
			parts[i] = p
		} else if strings.HasPrefix(p, "_, _ =") {
			rhs := strings.TrimSpace(strings.TrimPrefix(p, "_, _ ="))
			parts[i] = "_, _ = " + wrapBenchExpr(rhs)
		} else {
			parts[i] = wrapBenchExpr(p)
		}
	}
	return strings.Join(parts, "; ")
}

func wrapBenchExpr(expr string) string {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return expr
	}
	if strings.HasPrefix(expr, "solutions.") || strings.HasPrefix(expr, "(") {
		return expr
	}
	if i := strings.Index(expr, "{"); i > 0 {
		if j := strings.Index(expr[i:], "}"); j > 0 {
			end := i + j + 1
			if end < len(expr) && expr[end] == '.' {
				return "(solutions." + expr[:end] + ")" + expr[end:]
			}
		}
	}
	return "solutions." + expr
}

func benchFromExercise(slug, body string) string {
	if b, ok := benchBySlug[slug]; ok {
		return b
	}
	if name := firstExportedFunc(body); name != "" {
		return defaultBenchCall(name)
	}
	return `_ = 0`
}

func firstExportedFunc(body string) string {
	for _, line := range strings.Split(body, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "func ") {
			continue
		}
		rest := strings.TrimPrefix(line, "func ")
		if strings.HasPrefix(rest, "(") {
			continue
		}
		if idx := strings.Index(rest, "("); idx > 0 {
			return rest[:idx]
		}
	}
	return ""
}

func defaultBenchCall(fn string) string {
	calls := map[string]string{
		"MergeSort":        `a := []int{5, 2, 8, 1}; solutions.MergeSort(a)`,
		"FindMax":          `solutions.FindMax([]int{1, 5, 3})`,
		"AppendUnique":     `solutions.AppendUnique([]int{1}, 2)`,
		"MaxSumSubarrayK":  `solutions.MaxSumSubarrayK([]int{1, 2, 3, 4, 5}, 2)`,
		"BinarySearch":     `solutions.BinarySearch([]int{1, 2, 3}, 2)`,
		"GetUserEmail":     `_, _ = solutions.GetUserEmail(nil, "1")`,
		"Lines":            `solutions.Lines("a\nb\nc")`,
		"FizzBuzz":         `solutions.FizzBuzz(100)`,
		"NewStack":         `s := solutions.NewStack(); s.Push(1)`,
		"NewQueue":         `q := solutions.NewQueue(); q.Enqueue(1)`,
		"RuneLen":          `solutions.RuneLen("go")`,
		"Factorial":        `solutions.Factorial(5)`,
		"NewOrder":         `_, _ = solutions.NewOrder("o1", []string{"item"})`,
		"ParallelSum":      `solutions.ParallelSum([]int{1, 2, 3, 4}, 2)`,
		"CoinChange":       `solutions.CoinChange([]int{1, 2, 5}, 11)`,
		"CountOperations":  `solutions.CountOperations(1000)`,
		"TypeName":         `solutions.TypeName(42)`,
		"WordCount":        `solutions.WordCount([]string{"a", "b", "a"})`,
		"GCD":              `solutions.GCD(12, 18)`,
		"DailyTemperatures": `solutions.DailyTemperatures([]int{73, 74, 75})`,
	"Permute":          `solutions.Permute([]int{1, 2})`,
		"MaxActivities":    `solutions.MaxActivities([]int{1, 3, 2, 4})`,
		"RangeSum":         `solutions.RangeSum([]int{1, 2, 3})`,
		"NewUF":            `uf := solutions.NewUF(3); uf.Union(0, 1)`,
		"TreeHeight":       `solutions.TreeHeight(nil)`,
		"NewUserStore":     `s := solutions.NewUserStore(); s.Handle(solutions.CreateUserCommand{ID: "1", Email: "a@b.com"})`,
		"HasCycle":         `solutions.HasCycle(nil)`,
		"TwoSumSorted":     `solutions.TwoSumSorted([]int{1, 2, 3, 4}, 6)`,
		"SlidingWindowMax": `solutions.SlidingWindowMax([]int{1, 3, -1}, 2)`,
		"SingleNumber":     `solutions.SingleNumber([]int{1, 1, 2})`,
		"KMP":              `solutions.StrStr("ababc", "abc")`,
		"StrStr":           `solutions.StrStr("ababc", "abc")`,
		"InorderTraversal": `solutions.InorderTraversal(nil)`,
		"WithTimeoutOp":    `solutions.WithTimeoutOp(nil, 1)`,
		"SafeCounter":      `c := solutions.SafeCounter{}; c.Inc()`,
		"AtomicCounter":    `c := &solutions.AtomicCounter{}; c.Add(1)`,
		"FormatPrometheus": `solutions.FormatPrometheus("bench", 1)`,
		"Liveness":         `solutions.Liveness()`,
	}
	if c, ok := calls[fn]; ok {
		return c
	}
	return fmt.Sprintf(`solutions.%s()`, fn)
}

var benchBySlug = map[string]string{
	"06-pointers":               `a, b := 1, 2; solutions.Swap(&a, &b)`,
	"07-structs":                `(solutions.Rectangle{Width: 3, Height: 4}).Area()`,
	"08-methods":                `a := &solutions.BankAccount{}; a.Deposit(10)`,
	"19-stacks":                 `s := solutions.NewStack(); s.Push(1)`,
	"20-queues":                 `q := solutions.NewQueue(); q.Enqueue(1)`,
	"21-trees":                  `root := &solutions.TreeNode{Val: 1, Left: &solutions.TreeNode{Val: 2}}; solutions.TreeHeight(root)`,
	"22-binary-search-tree":     `root := solutions.InsertBST(nil, 5); solutions.SearchBST(root, 5)`,
	"28-union-find":           `uf := solutions.NewUF(3); uf.Union(0, 1)`,
	"29-segment-tree":         `solutions.RangeSum([]int{1, 2, 3})`,
	"30-fenwick-tree":         `solutions.RangeSum([]int{1, 2, 3})`,
	"38-backtracking":         `solutions.Permute([]int{1, 2})`,
	"40-greedy":               `solutions.MaxActivities([]int{1, 3, 2, 4})`,
	"45-graph-algorithms":     `g := solutions.Graph{1: {2}}; g.BFS(1)`,
	"46-tree-algorithms":      `solutions.TreeHeight(nil)`,
	"54-context":              `solutions.WithTimeoutOp(context.Background(), 1)`,
	"57-mutex":                `c := &solutions.SafeCounter{}; c.Inc()`,
	"60-atomic":               `c := &solutions.AtomicCounter{}; c.Add(1)`,
	"44-string-algorithms":    `solutions.StrStr("ababc", "abc")`,
	"113-cqrs":                `s := solutions.NewUserStore(); s.Handle(solutions.CreateUserCommand{ID: "1", Email: "a@b.com"})`,
	"114-event-driven":        `o := &solutions.OrderWithEvents{ID: "o1"}; o.Ship()`,
	"112-ddd":                   `_, _ = solutions.NewOrder("o1", []string{"item"})`,
	"26-trie":                   `tr := solutions.NewTrie(); tr.Insert("go")`,
	"27-graphs":                 `g := solutions.Graph{1: {2}}; g.BFS(1)`,
	"31-sorting":                `a := []int{5, 2, 8, 1}; solutions.MergeSort(a)`,
	"39-divide-conquer":         `a := []int{5, 2, 8, 1}; solutions.MergeSort(a)`,
	"49-goroutines":             `solutions.ParallelSum([]int{1, 2, 3, 4}, 2)`,
	"147-production-projects":   `(solutions.ModuleProgress{Module: "b", Percent: 100}).Completed()`,
	"148-interview-preparation": `(solutions.ModuleProgress{Module: "b", Percent: 100}).Completed()`,
	"149-cheat-sheets":          `(solutions.ModuleProgress{Module: "b", Percent: 100}).Completed()`,
	"150-roadmap":               `(solutions.ModuleProgress{Module: "b", Percent: 100}).Completed()`,
}
