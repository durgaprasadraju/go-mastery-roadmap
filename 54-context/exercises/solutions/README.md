# Context — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Context** concurrency patterns in Go.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Context exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Propagate context cancellation |

## Run

```bash
go test ./54-context/exercises/solutions/
go test -bench=. ./54-context/exercises/solutions/
```
