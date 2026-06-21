# Control Flow — Exercises

Hands-on exercises for **Control Flow** — variables, types, idioms, and Go fundamentals.

## Exercise 1: Core Implementation

Implement and understand the core **Control Flow** concept in exercise1.go (try yourself first).

**Run tests:**
```bash
go test ./03-control-flow/exercises/solutions/ -run TestExercise1
```

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

```bash
go test ./03-control-flow/exercises/solutions/ -run TestService
```

## Exercise 3: Benchmarks

Measure performance of the core implementation.

```bash
go test -bench=. ./03-control-flow/exercises/solutions/
```

## Exercise 4: Interview Challenge

**Problem:** FizzBuzz(n) returning slice of strings 1..n with Fizz/Buzz rules

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

```bash
go test ./03-control-flow/exercises/solutions/ -run TestInterviewChallenge
```

> Solutions are in solutions/ — attempt all exercises first.
