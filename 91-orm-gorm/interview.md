# GORM ORM — Interview Questions

## Beginner

### Q1: What is GORM ORM and why does it matter in Go?

**Answer:** GORM ORM is a core Go concept used in every production codebase.

**Deep Explanation:** Understanding GORM ORM demonstrates language fundamentals and informs design decisions in concurrent, distributed systems.

**Follow-up Questions:**
- How would you explain this to a junior developer?
- What Go standard library packages relate to this?

**Production Scenario:** You are building a high-throughput API. How does GORM ORM affect latency, memory, and correctness?

---

## Intermediate

### Q2: What are common pitfalls with GORM ORM?

**Answer:** See `common-mistakes.md` for detailed pitfalls.

**Deep Explanation:** Most production bugs stem from misunderstanding ownership, concurrency, or error propagation.

**Follow-up Questions:**
- How do you test edge cases?
- What benchmarks would you write?

**Production Scenario:** A service shows elevated p99 latency. How does GORM ORM factor into your investigation?

---

## Advanced

### Q3: How would you design a system leveraging GORM ORM at scale?

**Answer:** Apply Clean Architecture layers, context propagation, and observability.

**Deep Explanation:** Staff-level answers cover trade-offs: consistency vs availability, coupling vs cohesion, operational complexity.

**Follow-up Questions:**
- How does this interact with the Go memory model?
- What metrics and alerts would you add?

**Production Scenario:** Design for 10x traffic growth. What changes to your GORM ORM approach?

---

## Staff Engineer

### Q4: What architectural decisions involve GORM ORM across a microservices platform?

**Answer:** Cross-cutting concerns include shared libraries, versioning, backward compatibility, and team boundaries.

**Deep Explanation:** Staff engineers evaluate org-wide impact: developer velocity, incident frequency, and total cost of ownership.

**Follow-up Questions:**
- How do you prevent anti-patterns from spreading?
- What would you standardize in a platform team?

**Production Scenario:** Multi-region deployment with strict SLAs — how does GORM ORM influence your design?
