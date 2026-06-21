# Terraform — Exercise Solutions

> Attempt exercises in ../README.md before reading these solutions.

## Topic

**Terraform** — containers, K8s, CI/CD, IaC.

## Files

| File | Description |
|------|-------------|
| exercise1.go | Core Terraform exercise |
| exercise1_test.go | Unit tests |
| exercise2.go | Production patterns (context + structured errors) |
| exercise3_bench_test.go | Performance benchmarks |
| interview-challenge-solution.go | Design Kubernetes liveness/readiness probes |

## Run

```bash
go test ./130-terraform/exercises/solutions/
go test -bench=. ./130-terraform/exercises/solutions/
```
