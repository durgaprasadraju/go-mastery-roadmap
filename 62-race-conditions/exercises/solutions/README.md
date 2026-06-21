# Race Conditions — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Race Conditions** — goroutines, channels, sync primitives.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Race Conditions exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Coordinate concurrent work safely |

## Run

```bash
go test ./62-race-conditions/exercises/solutions/
go test -bench=. ./62-race-conditions/exercises/solutions/
```
