# Server-Sent Events — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Server-Sent Events** — protocols, clients, servers.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Server-Sent Events exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design HTTP handler with middleware chain |

## Run

```bash
go test ./80-server-sent-events/exercises/solutions/
go test -bench=. ./80-server-sent-events/exercises/solutions/
```
