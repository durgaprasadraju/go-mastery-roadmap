# Grafana — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Grafana** — logs, metrics, traces.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Grafana exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design RED/USE metrics for a service |

## Run

```bash
go test ./122-grafana/exercises/solutions/
go test -bench=. ./122-grafana/exercises/solutions/
```
