# Fan-In / Fan-Out — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Fan-In / Fan-Out** — goroutines, channels, sync primitives.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Fan-In / Fan-Out exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Coordinate concurrent work safely |

## Run

```bash
go test ./53-fan-in-fan-out/exercises/solutions/
go test -bench=. ./53-fan-in-fan-out/exercises/solutions/
```
