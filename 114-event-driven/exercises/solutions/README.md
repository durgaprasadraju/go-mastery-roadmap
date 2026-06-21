# Event-Driven Architecture — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Event-Driven Architecture** — production architecture patterns.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Event-Driven Architecture exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Publish domain events from aggregate |

## Run

```bash
go test ./114-event-driven/exercises/solutions/
go test -bench=. ./114-event-driven/exercises/solutions/
```
