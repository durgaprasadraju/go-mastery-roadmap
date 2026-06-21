# RabbitMQ — Exercises

**RabbitMQ** — pub/sub, streaming, queues.

## Exercise 1: Core Implementation

Implement and understand the core **RabbitMQ** concept in exercise1.go (try yourself first).

**Run tests:**
```bash
go test ./133-rabbitmq/exercises/solutions/ -run TestExercise1
```

## Exercise 2: Production Patterns

Add context cancellation and structured errors in your own version, then compare with exercise2.go.

```bash
go test ./133-rabbitmq/exercises/solutions/ -run TestService
```

## Exercise 3: Benchmarks

Measure performance of the core implementation.

```bash
go test -bench=. ./133-rabbitmq/exercises/solutions/
```

## Exercise 4: Interview Challenge

**Problem:** Design idempotent Kafka consumer

Work in interview-challenge.go for 30 minutes before checking solutions/interview-challenge-solution.go.

```bash
go test ./133-rabbitmq/exercises/solutions/ -run TestInterviewChallenge
```

> Solutions are in solutions/ — attempt all exercises first.
