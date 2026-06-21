# Load Balancing — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Load Balancing** — protocols, clients, servers.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Load Balancing exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design HTTP handler with middleware chain |

## Run

```bash
go test ./82-load-balancing/exercises/solutions/
go test -bench=. ./82-load-balancing/exercises/solutions/
```
