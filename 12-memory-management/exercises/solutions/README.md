# Memory Management — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

Hands-on exercises for **Memory Management** — variables, types, idioms, and Go fundamentals.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Memory Management exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Detect if slice reuses same backing array after append |

## Run

```bash
go test ./12-memory-management/exercises/solutions/
go test -bench=. ./12-memory-management/exercises/solutions/
```
