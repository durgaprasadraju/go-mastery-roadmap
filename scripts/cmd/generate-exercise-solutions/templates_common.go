package main

import (
	"fmt"
	"strings"
)

func readme(m module, s topicSpec) string {
	return fmt.Sprintf(`# %s — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

%s

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core %s exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | %s |

## Run

`+"```bash\n"+`go test ./%s/exercises/solutions/
go test -bench=. ./%s/exercises/solutions/
`+"```"+`
`, m.title, s.summary, m.title, s.interviewProblem, m.slug, m.slug)
}

func exercise2(m module, s topicSpec) string {
	return fmt.Sprintf(`package solutions

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
)

// Service demonstrates production patterns for %s.
type Service struct {
	Name string
}

// Process handles %s operations with context cancellation.
func (s *Service) Process(ctx context.Context, input string) (string, error) {
	select {
	case <-ctx.Done():
		return "", apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "%s cancelled")
	default:
	}
	if input == "" {
		return "", apperrors.New(apperrors.CodeValidation, "input required")
	}
	return fmt.Sprintf("[%%s] %%s: %%s", s.Name, %q, input), nil
}
`, m.title, m.title, m.title, m.title)
}

func exercise2Test(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"context"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestServiceProcess(t *testing.T) {
	svc := &solutions.Service{Name: "%s"}
	out, err := svc.Process(context.Background(), "test")
	if err != nil || out == "" {
		t.Fatalf("out=%%q err=%%v", out, err)
	}
}

func TestServiceValidation(t *testing.T) {
	svc := &solutions.Service{Name: "%s"}
	_, err := svc.Process(context.Background(), "")
	if err == nil {
		t.Fatal("expected validation error")
	}
}
`, m.slug, m.title, m.title)
}

func exercise3Bench(m module, s topicSpec) string {
	call := s.benchCall
	extraImport := ""
	if strings.Contains(call, "context.") {
		extraImport = "\n\t\"context\""
	}
	return fmt.Sprintf(`package solutions_test

import (
	"testing"%s
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func BenchmarkExercise1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		%s
	}
}
`, extraImport, m.slug, call)
}

func exercisesReadmeContent(m module, s topicSpec) string {
	return fmt.Sprintf(`# %s — Exercises

%s

## Exercise 1: Core Implementation

Implement and understand the core **%s** concept in exercise1.go (try yourself first).

**Run tests:**
`+"```bash\n"+`go test ./%s/exercises/solutions/ -run TestExercise1
`+"```"+`

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

`+"```bash\n"+`go test ./%s/exercises/solutions/ -run TestService
`+"```"+`

## Exercise 3: Benchmarks

Measure performance of the core implementation.

`+"```bash\n"+`go test -bench=. ./%s/exercises/solutions/
`+"```"+`

## Exercise 4: Interview Challenge

**Problem:** %s

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

`+"```bash\n"+`go test ./%s/exercises/solutions/ -run TestInterviewChallenge
`+"```"+`

> Solutions are in solutions/ — attempt all exercises first.
`, m.title, s.summary, m.title, m.slug, m.slug, m.slug, s.interviewProblem, m.slug)
}

func interviewChallengeStub(m module, s topicSpec) string {
	return fmt.Sprintf(`// Package exercises — solve before viewing solutions/interview-challenge-solution.go
package exercises

// InterviewChallenge: %s
// Requirements: correct complexity, edge cases, tests.
func InterviewChallenge() {
	// TODO: implement for %s
}
`, s.interviewProblem, m.title)
}
