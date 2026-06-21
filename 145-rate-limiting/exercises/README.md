# Rate Limiting — Exercises

**Rate Limiting** — scalable system components.

## Exercise 1: Core Implementation

Implement and understand the core **Rate Limiting** concept in exercise1.go (try yourself first).

**Run tests:**
```bash
go test ./145-rate-limiting/exercises/solutions/ -run TestExercise1
```

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

```bash
go test ./145-rate-limiting/exercises/solutions/ -run TestService
```

## Exercise 3: Benchmarks

Measure performance of the core implementation.

```bash
go test -bench=. ./145-rate-limiting/exercises/solutions/
```

## Exercise 4: Interview Challenge

**Problem:** Token bucket rate limiter

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

```bash
go test ./145-rate-limiting/exercises/solutions/ -run TestInterviewChallenge
```

> Solutions are in solutions/ — attempt all exercises first.
