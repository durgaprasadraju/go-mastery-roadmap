# Apache Kafka — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Apache Kafka** — pub/sub, streaming, queues.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Apache Kafka exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design idempotent Kafka consumer |

## Run

```bash
go test ./132-kafka/exercises/solutions/
go test -bench=. ./132-kafka/exercises/solutions/
```
