# Q159: Goroutines

## Answer

This is a **intermediate**-level Go interview question covering **Goroutines**.

## Deep Explanation

Goroutines is a critical concept for Go developers at the intermediate level. Refer to the corresponding roadmap module for theory, diagrams, and production code examples.

**Key concepts:**
- Fundamental theory and mental models
- Idiomatic Go implementation patterns
- Performance characteristics and complexity analysis
- Common pitfalls and how to avoid them in production

```go
// Illustrative pattern — see module examples for full implementation
func Example() error {
    // Implementation specific to Goroutines
    return nil
}
```

## Follow-up Questions

1. How would you unit test code involving Goroutines?
2. What are the performance implications of Goroutines at scale?
3. How does Goroutines interact with Go's concurrency model and memory model?
4. What metrics would you monitor in production?

## Production Scenario

Your team reports elevated p99 latency in a service where Goroutines is heavily used. Walk through your investigation:

1. Which metrics and logs do you check first?
2. How do you reproduce the issue locally?
3. What is your fix and rollout strategy?
4. How do you prevent regression?
