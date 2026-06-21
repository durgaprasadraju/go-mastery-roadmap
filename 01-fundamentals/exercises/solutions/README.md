# Go Fundamentals — Exercise Solutions

> Attempt the exercises in `../README.md` before reading these solutions.

## Files

| File | Exercise | Description |
|------|----------|-------------|
| `exercise1.go` | Exercise 1 | Core Go Fundamentals implementation |
| `exercise1_test.go` | Exercise 1 | Unit tests |
| `exercise2.go` | Exercise 2 | Production patterns (context, errors) |
| `exercise2_test.go` | Exercise 2 | Tests for success, validation, cancellation |
| `exercise3_bench_test.go` | Exercise 3 | Benchmarks and complexity notes |
| `interview-challenge-solution.go` | Exercise 4 | Kadane's algorithm (max subarray sum) |
| `interview-challenge-solution_test.go` | Exercise 4 | Interview challenge test |

See `../README.md` for:

- Exercise 2 — context cancellation and failure code examples
- Exercise 3 — benchmark line-by-line guide (`b.ResetTimer`, `b.N`, etc.)
- Exercise 4 — interview challenge problem, solution walkthrough, edge cases, and **13 follow-up Q&As**

## Run Tests

```bash
go test ./01-fundamentals/exercises/solutions/...
```

## Run Benchmarks

```bash
go test -bench=. ./01-fundamentals/exercises/solutions/...
go test -bench=. -benchmem ./01-fundamentals/exercises/solutions/...
```
