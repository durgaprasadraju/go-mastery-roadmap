# Redis — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Redis** — pub/sub, streaming, queues.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Redis exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design idempotent Kafka consumer |

## Run

```bash
go test ./131-redis/exercises/solutions/
go test -bench=. ./131-redis/exercises/solutions/
```
