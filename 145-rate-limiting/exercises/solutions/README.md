# Rate Limiting — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Rate Limiting** — scalable system components.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Rate Limiting exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Token bucket rate limiter |

## Run

```bash
go test ./145-rate-limiting/exercises/solutions/
go test -bench=. ./145-rate-limiting/exercises/solutions/
```
