# RabbitMQ — Common Mistakes

## 1. Ignoring Error Returns

Go functions return errors for a reason. Never use `_ = err` in production code without explicit justification.

## 2. Missing Context Cancellation

Long-running operations without `context.Context` cannot be cancelled during shutdown or timeout.

## 3. Premature Optimization

Optimize based on benchmarks and profiles, not assumptions. Use `pprof` and `benchstat`.

## 4. Shared Mutable State Without Synchronization

Concurrent access to shared data requires mutexes, channels, or atomic operations. Run `go test -race`.

## 5. Interface Pollution

Define interfaces at the consumer, not the producer. Keep interfaces small (1-3 methods).

## 6. Not Using `defer` for Cleanup

Always `defer` file closes, mutex unlocks, and connection releases to prevent leaks on panic paths.

## 7. Logging Instead of Returning Errors

Log at system boundaries (HTTP handlers, main). Libraries should return errors; callers decide logging.

## 8. Hardcoding Configuration

Use environment variables and config structs. Support 12-factor app principles.
