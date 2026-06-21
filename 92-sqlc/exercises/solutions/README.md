# sqlc — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**sqlc** — SQL, pooling, transactions.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core sqlc exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design repository pattern with transactions |

## Run

```bash
go test ./92-sqlc/exercises/solutions/
go test -bench=. ./92-sqlc/exercises/solutions/
```
