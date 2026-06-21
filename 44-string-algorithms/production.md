# String Algorithms — Production Considerations

## Structured Logging

```go
log.Info("operation completed",
    zap.String("module", "String Algorithms"),
    zap.Duration("elapsed", elapsed),
)
```

## Context Propagation

Always pass `context.Context` as the first parameter. Use it for cancellation, deadlines, and trace IDs.

## Graceful Shutdown

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
// drain connections, flush buffers
```

## Configuration

Use environment variables via `pkg/config`. Never hardcode secrets.

## Health Checks

Register dependency checkers with `pkg/health` for Kubernetes readiness probes.

## Metrics & Tracing

- Expose Prometheus metrics for latency, throughput, and error rates
- Propagate OpenTelemetry trace context across service boundaries

## Error Handling

Wrap errors with `pkg/apperrors`. Return structured errors to callers; log at boundaries.

## Security

- Validate all inputs
- Use parameterized queries for SQL
- Apply least-privilege for service accounts
- Rotate secrets via vault/Kubernetes secrets

## Deployment Checklist

- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Benchmarks within SLA
- [ ] No race conditions (`go test -race`)
- [ ] Liveness/readiness probes configured
- [ ] Dashboards and alerts in place
