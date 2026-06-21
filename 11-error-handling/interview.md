# Error Handling — Interview Questions

## Beginner

### Q1: What is Error Handling and why does it matter in Go?

**Answer:** Error Handling is a core Go concept used in every production codebase.

**Deep Explanation:** Understanding Error Handling demonstrates language fundamentals and informs design decisions in concurrent, distributed systems.

**Follow-up Questions:**
- How would you explain this to a junior developer?
- What Go standard library packages relate to this?

**Production Scenario:** You are building a high-throughput API. How does Error Handling affect latency, memory, and correctness?

---

## Intermediate

### Q2: What are common pitfalls with Error Handling?

**Answer:** See `common-mistakes.md` for detailed pitfalls.

**Deep Explanation:** Most production bugs stem from misunderstanding ownership, concurrency, or error propagation.

**Follow-up Questions:**
- How do you test edge cases?
- What benchmarks would you write?

**Production Scenario:** A service shows elevated p99 latency. How does Error Handling factor into your investigation?

---

## Advanced

### Q3: How would you design a system leveraging Error Handling at scale?

**Answer:** Apply Clean Architecture layers, context propagation, and observability.

**Deep Explanation:** Staff-level answers cover trade-offs: consistency vs availability, coupling vs cohesion, operational complexity.

**Follow-up Questions:**
- How does this interact with the Go memory model?
- What metrics and alerts would you add?

**Production Scenario:** Design for 10x traffic growth. What changes to your Error Handling approach?

---

## Staff Engineer

### Q4: What architectural decisions involve Error Handling across a microservices platform?

**Answer:** Cross-cutting concerns include shared libraries, versioning, backward compatibility, and team boundaries.

**Deep Explanation:** Staff engineers evaluate org-wide impact: developer velocity, incident frequency, and total cost of ownership.

**Follow-up Questions:**
- How do you prevent anti-patterns from spreading?
- What would you standardize in a platform team?

**Production Scenario:** Multi-region deployment with strict SLAs — how does Error Handling influence your design?
