# Observability — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Observability** — logs, metrics, traces.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Observability exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design RED/USE metrics for a service |

## Run

```bash
go test ./120-observability/exercises/solutions/
go test -bench=. ./120-observability/exercises/solutions/
```
