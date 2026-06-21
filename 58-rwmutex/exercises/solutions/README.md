# RWMutex — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**RWMutex** — goroutines, channels, sync primitives.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core RWMutex exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Coordinate concurrent work safely |

## Run

```bash
go test ./58-rwmutex/exercises/solutions/
go test -bench=. ./58-rwmutex/exercises/solutions/
```
