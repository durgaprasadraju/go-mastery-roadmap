# Repository Pattern — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Repository Pattern** — production architecture patterns.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Repository Pattern exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Generic repository interface with in-memory impl |

## Run

```bash
go test ./118-repository-pattern/exercises/solutions/
go test -bench=. ./118-repository-pattern/exercises/solutions/
```
