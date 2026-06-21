# Atomic Operations — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Atomic Operations** concurrency patterns in Go.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Atomic Operations exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Lock-free counter with atomic.Int64 |

## Run

```bash
go test ./60-atomic/exercises/solutions/
go test -bench=. ./60-atomic/exercises/solutions/
```
