# Kubernetes — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Kubernetes** — containers, K8s, CI/CD, IaC.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Kubernetes exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design Kubernetes liveness/readiness probes |

## Run

```bash
go test ./126-kubernetes/exercises/solutions/
go test -bench=. ./126-kubernetes/exercises/solutions/
```
