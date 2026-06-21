# Circuit Breaker — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Circuit Breaker** — scalable system components.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Circuit Breaker exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Circuit breaker with open/half-open/closed |

## Run

```bash
go test ./146-circuit-breaker/exercises/solutions/
go test -bench=. ./146-circuit-breaker/exercises/solutions/
```
