# PostgreSQL — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**PostgreSQL** — SQL, pooling, transactions.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core PostgreSQL exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design repository pattern with transactions |

## Run

```bash
go test ./84-postgresql/exercises/solutions/
go test -bench=. ./84-postgresql/exercises/solutions/
```
