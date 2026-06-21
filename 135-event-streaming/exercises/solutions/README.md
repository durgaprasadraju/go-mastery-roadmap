# Event Streaming — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Event Streaming** — pub/sub, streaming, queues.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Event Streaming exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design idempotent Kafka consumer |

## Run

```bash
go test ./135-event-streaming/exercises/solutions/
go test -bench=. ./135-event-streaming/exercises/solutions/
```
