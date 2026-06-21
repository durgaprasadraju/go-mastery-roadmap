# OpenTelemetry — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**OpenTelemetry** — logs, metrics, traces.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core OpenTelemetry exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design RED/USE metrics for a service |

## Run

```bash
go test ./123-opentelemetry/exercises/solutions/
go test -bench=. ./123-opentelemetry/exercises/solutions/
```
