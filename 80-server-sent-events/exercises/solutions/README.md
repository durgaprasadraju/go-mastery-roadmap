# Server-Sent Events — Exercise Solutions

> Attempt the exercises in `../README.md` before reading these solutions.

## Files

| File | Exercise | Description |
|------|----------|-------------|
| `exercise1.go` | Exercise 1 | Core Server-Sent Events implementation |
| `exercise1_test.go` | Exercise 1 | Unit tests |
| `exercise2.go` | Exercise 2 | Production patterns (context, errors) |
| `exercise3_bench_test.go` | Exercise 3 | Benchmarks and complexity notes |

## Run Tests

```bash
go test ./80-server-sent-events/exercises/solutions/...
```

## Run Benchmarks

```bash
go test -bench=. ./80-server-sent-events/exercises/solutions/...
```
