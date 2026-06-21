# sqlc — Exercises

**sqlc** — SQL, pooling, transactions.

## Exercise 1: Core Implementation

Implement and understand the core **sqlc** concept in exercise1.go (try yourself first).

**Run tests:**
```bash
go test ./92-sqlc/exercises/solutions/ -run TestExercise1
```

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

```bash
go test ./92-sqlc/exercises/solutions/ -run TestService
```

## Exercise 3: Benchmarks

Measure performance of the core implementation.

```bash
go test -bench=. ./92-sqlc/exercises/solutions/
```

## Exercise 4: Interview Challenge

**Problem:** Design repository pattern with transactions

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

```bash
go test ./92-sqlc/exercises/solutions/ -run TestInterviewChallenge
```

> Solutions are in solutions/ — attempt all exercises first.
