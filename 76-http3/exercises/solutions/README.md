# HTTP/3 — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**HTTP/3** — protocols, clients, servers.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core HTTP/3 exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design HTTP handler with middleware chain |

## Run

```bash
go test ./76-http3/exercises/solutions/
go test -bench=. ./76-http3/exercises/solutions/
```
