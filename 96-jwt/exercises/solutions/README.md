# JWT — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**JWT** — create, parse, and validate JSON Web Tokens.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core JWT exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Validate JWT structure and expiry |

## Run

```bash
go test ./96-jwt/exercises/solutions/
go test -bench=. ./96-jwt/exercises/solutions/
```
