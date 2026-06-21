# HTTP/2 — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**HTTP/2** — protocols, clients, servers.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core HTTP/2 exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design HTTP handler with middleware chain |

## Run

```bash
go test ./75-http2/exercises/solutions/
go test -bench=. ./75-http2/exercises/solutions/
```
