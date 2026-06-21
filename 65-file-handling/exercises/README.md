# File Handling — Exercises

**File Handling** — file I/O and serialization.

## Exercise 1: Core Implementation

Implement and understand the core **File Handling** concept in exercise1.go (try yourself first).

**Run tests:**
```bash
go test ./65-file-handling/exercises/solutions/ -run TestExercise1
```

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

```bash
go test ./65-file-handling/exercises/solutions/ -run TestService
```

## Exercise 3: Benchmarks

Measure performance of the core implementation.

```bash
go test -bench=. ./65-file-handling/exercises/solutions/
```

## Exercise 4: Interview Challenge

**Problem:** Stream-parse large JSON without loading entire file

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

```bash
go test ./65-file-handling/exercises/solutions/ -run TestInterviewChallenge
```

> Solutions are in solutions/ — attempt all exercises first.
