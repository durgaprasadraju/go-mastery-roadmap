# Clean Architecture — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Clean Architecture** — production architecture patterns.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Clean Architecture exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Separate handler, use case, and repository layers |

## Run

```bash
go test ./111-clean-architecture/exercises/solutions/
go test -bench=. ./111-clean-architecture/exercises/solutions/
```
