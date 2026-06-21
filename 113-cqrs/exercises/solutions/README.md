# CQRS — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**CQRS** — production architecture patterns.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core CQRS exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Separate command and query handlers |

## Run

```bash
go test ./113-cqrs/exercises/solutions/
go test -bench=. ./113-cqrs/exercises/solutions/
```
