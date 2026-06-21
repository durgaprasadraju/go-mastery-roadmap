# Packages & Modules — Exercises

Hands-on exercises for **Packages & Modules** — variables, types, idioms, and Go fundamentals.

## Exercise 1: Core Implementation

Implement and understand the core **Packages & Modules** concept in exercise1.go (try yourself first).

**Run tests:**
```bash
go test ./05-packages-modules/exercises/solutions/ -run TestExercise1
```

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

```bash
go test ./05-packages-modules/exercises/solutions/ -run TestService
```

## Exercise 3: Benchmarks

Measure performance of the core implementation.

```bash
go test -bench=. ./05-packages-modules/exercises/solutions/
```

## Exercise 4: Interview Challenge

**Problem:** ParseSemVer parses major.minor.patch from string

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

```bash
go test ./05-packages-modules/exercises/solutions/ -run TestInterviewChallenge
```

> Solutions are in solutions/ — attempt all exercises first.
