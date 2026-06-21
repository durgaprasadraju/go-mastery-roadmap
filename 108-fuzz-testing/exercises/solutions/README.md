# Fuzz Testing — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Fuzz Testing** — unit, integration, fuzz, benchmarks.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Fuzz Testing exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Write table-driven tests with subtests |

## Run

```bash
go test ./108-fuzz-testing/exercises/solutions/
go test -bench=. ./108-fuzz-testing/exercises/solutions/
```
