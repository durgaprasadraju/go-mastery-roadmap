#!/usr/bin/env bash
# generate-modules.sh — Bootstrap all 150 learning modules with consistent structure.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

# All 150 modules: "NN-slug|Title|Category|Level"
MODULES=(
"01-fundamentals|Go Fundamentals|fundamentals|beginner"
"02-types-system|Go Type System|fundamentals|beginner"
"03-control-flow|Control Flow|fundamentals|beginner"
"04-functions|Functions & Closures|fundamentals|beginner"
"05-packages-modules|Packages & Modules|fundamentals|beginner"
"06-pointers|Pointers|fundamentals|beginner"
"07-structs|Structs|fundamentals|beginner"
"08-methods|Methods|fundamentals|beginner"
"09-interfaces|Interfaces|fundamentals|intermediate"
"10-generics|Generics|fundamentals|intermediate"
"11-error-handling|Error Handling|fundamentals|intermediate"
"12-memory-management|Memory Management|fundamentals|intermediate"
"13-arrays|Arrays|data-structures|beginner"
"14-slices|Slices|data-structures|beginner"
"15-maps|Maps|data-structures|beginner"
"16-strings-runes-bytes|Strings, Runes & Bytes|data-structures|beginner"
"17-recursion|Recursion|data-structures|beginner"
"18-linked-lists|Linked Lists|data-structures|intermediate"
"19-stacks|Stacks|data-structures|intermediate"
"20-queues|Queues|data-structures|intermediate"
"21-trees|Trees|data-structures|intermediate"
"22-binary-search-tree|Binary Search Trees|data-structures|intermediate"
"23-heaps|Heaps|data-structures|intermediate"
"24-priority-queue|Priority Queues|data-structures|intermediate"
"25-hash-tables|Hash Tables|data-structures|intermediate"
"26-trie|Tries|data-structures|intermediate"
"27-graphs|Graphs|data-structures|intermediate"
"28-union-find|Union-Find|data-structures|advanced"
"29-segment-tree|Segment Trees|data-structures|advanced"
"30-fenwick-tree|Fenwick Trees|data-structures|advanced"
"31-sorting|Sorting Algorithms|algorithms|intermediate"
"32-searching|Searching Algorithms|algorithms|intermediate"
"33-sliding-window|Sliding Window|algorithms|intermediate"
"34-two-pointers|Two Pointers|algorithms|intermediate"
"35-fast-slow-pointers|Fast & Slow Pointers|algorithms|intermediate"
"36-monotonic-stack|Monotonic Stack|algorithms|intermediate"
"37-monotonic-queue|Monotonic Queue|algorithms|intermediate"
"38-backtracking|Backtracking|algorithms|advanced"
"39-divide-conquer|Divide & Conquer|algorithms|advanced"
"40-greedy|Greedy Algorithms|algorithms|intermediate"
"41-dynamic-programming|Dynamic Programming|algorithms|advanced"
"42-bit-manipulation|Bit Manipulation|algorithms|intermediate"
"43-math-algorithms|Math Algorithms|algorithms|intermediate"
"44-string-algorithms|String Algorithms|algorithms|advanced"
"45-graph-algorithms|Graph Algorithms|algorithms|advanced"
"46-tree-algorithms|Tree Algorithms|algorithms|advanced"
"47-big-o-analysis|Big-O Analysis|complexity|beginner"
"48-complexity-analysis|Complexity Analysis|complexity|intermediate"
"49-goroutines|Goroutines|concurrency|intermediate"
"50-channels|Channels|concurrency|intermediate"
"51-worker-pools|Worker Pools|concurrency|intermediate"
"52-pipelines|Pipelines|concurrency|intermediate"
"53-fan-in-fan-out|Fan-In / Fan-Out|concurrency|advanced"
"54-context|Context|concurrency|intermediate"
"55-select|Select Statement|concurrency|intermediate"
"56-sync-package|sync Package|concurrency|intermediate"
"57-mutex|Mutex|concurrency|intermediate"
"58-rwmutex|RWMutex|concurrency|intermediate"
"59-cond|sync.Cond|concurrency|advanced"
"60-atomic|Atomic Operations|concurrency|advanced"
"61-memory-model|Go Memory Model|concurrency|advanced"
"62-race-conditions|Race Conditions|concurrency|advanced"
"63-deadlocks|Deadlocks|concurrency|advanced"
"64-concurrency-patterns|Concurrency Patterns|concurrency|advanced"
"65-file-handling|File Handling|io|beginner"
"66-json|JSON|io|beginner"
"67-yaml|YAML|io|beginner"
"68-csv|CSV|io|beginner"
"69-xml|XML|io|beginner"
"70-compression|Compression|io|intermediate"
"71-networking-fundamentals|Networking Fundamentals|networking|beginner"
"72-tcp|TCP|networking|intermediate"
"73-udp|UDP|networking|intermediate"
"74-http|HTTP|networking|intermediate"
"75-http2|HTTP/2|networking|advanced"
"76-http3|HTTP/3|networking|advanced"
"77-rest-api|REST API|networking|intermediate"
"78-grpc|gRPC|networking|advanced"
"79-websocket|WebSocket|networking|intermediate"
"80-server-sent-events|Server-Sent Events|networking|intermediate"
"81-dns|DNS|networking|intermediate"
"82-load-balancing|Load Balancing|networking|advanced"
"83-database-sql|SQL Fundamentals|database|beginner"
"84-postgresql|PostgreSQL|database|intermediate"
"85-mysql|MySQL|database|intermediate"
"86-sqlite|SQLite|database|beginner"
"87-transactions|Transactions|database|intermediate"
"88-indexing|Indexing|database|intermediate"
"89-query-optimization|Query Optimization|database|advanced"
"90-connection-pooling|Connection Pooling|database|intermediate"
"91-orm-gorm|GORM ORM|database|intermediate"
"92-sqlc|sqlc|database|intermediate"
"93-migrations|Database Migrations|database|intermediate"
"94-authentication|Authentication|security|intermediate"
"95-authorization|Authorization|security|intermediate"
"96-jwt|JWT|security|intermediate"
"97-oauth2|OAuth2|security|advanced"
"98-session-management|Session Management|security|intermediate"
"99-security|Security Fundamentals|security|intermediate"
"100-encryption|Encryption|security|advanced"
"101-hashing|Hashing|security|intermediate"
"102-secrets-management|Secrets Management|security|advanced"
"103-testing|Testing Fundamentals|testing|beginner"
"104-unit-testing|Unit Testing|testing|beginner"
"105-integration-testing|Integration Testing|testing|intermediate"
"106-e2e-testing|E2E Testing|testing|intermediate"
"107-mocking|Mocking|testing|intermediate"
"108-fuzz-testing|Fuzz Testing|testing|advanced"
"109-benchmarking|Benchmarking|testing|intermediate"
"110-profiling|Profiling|testing|advanced"
"111-clean-architecture|Clean Architecture|architecture|advanced"
"112-ddd|Domain-Driven Design|architecture|advanced"
"113-cqrs|CQRS|architecture|advanced"
"114-event-driven|Event-Driven Architecture|architecture|advanced"
"115-microservices|Microservices|architecture|advanced"
"116-monoliths|Monoliths|architecture|intermediate"
"117-hexagonal|Hexagonal Architecture|architecture|advanced"
"118-repository-pattern|Repository Pattern|architecture|intermediate"
"119-logging|Logging|observability|beginner"
"120-observability|Observability|observability|intermediate"
"121-prometheus|Prometheus|observability|intermediate"
"122-grafana|Grafana|observability|intermediate"
"123-opentelemetry|OpenTelemetry|observability|advanced"
"124-tracing|Distributed Tracing|observability|advanced"
"125-docker|Docker|cloud-native|intermediate"
"126-kubernetes|Kubernetes|cloud-native|advanced"
"127-helm|Helm|cloud-native|advanced"
"128-ci-cd|CI/CD|cloud-native|intermediate"
"129-github-actions|GitHub Actions|cloud-native|intermediate"
"130-terraform|Terraform|cloud-native|advanced"
"131-redis|Redis|messaging|intermediate"
"132-kafka|Apache Kafka|messaging|advanced"
"133-rabbitmq|RabbitMQ|messaging|intermediate"
"134-nats|NATS|messaging|advanced"
"135-event-streaming|Event Streaming|messaging|advanced"
"136-design-patterns|Design Patterns|patterns|advanced"
"137-system-design|System Design|system-design|advanced"
"138-distributed-systems|Distributed Systems|system-design|advanced"
"139-consensus|Consensus Algorithms|system-design|advanced"
"140-raft|Raft Consensus|system-design|advanced"
"141-cap-theorem|CAP Theorem|system-design|advanced"
"142-sharding|Sharding|system-design|advanced"
"143-replication|Replication|system-design|advanced"
"144-caching|Caching Strategies|system-design|advanced"
"145-rate-limiting|Rate Limiting|system-design|advanced"
"146-circuit-breaker|Circuit Breaker|system-design|advanced"
"147-production-projects|Production Projects|projects|advanced"
"148-interview-preparation|Interview Preparation|interview|all-levels"
"149-cheat-sheets|Cheat Sheets|reference|all-levels"
"150-roadmap|Learning Roadmap|reference|all-levels"
)

is_ds_module() {
  local slug="$1"
  [[ "$slug" =~ ^(13|14|15|16|17|18|19|20|21|22|23|24|25|26|27|28|29|30)- ]]
}

is_algo_module() {
  local slug="$1"
  [[ "$slug" =~ ^(31|32|33|34|35|36|37|38|39|40|41|42|43|44|45|46)- ]]
}

create_readme() {
  local dir="$1" slug="$2" title="$3" category="$4" level="$5"
  cat > "$dir/README.md" << EOF
# ${title}

> **Category:** ${category} | **Level:** ${level} | **Module:** \`${slug}\`

## Overview

This module covers **${title}** — a core topic on the Go Mastery Roadmap from beginner to staff engineer level.

## Learning Objectives

- Understand the theory and mental models behind ${title}
- Implement production-grade Go code following Clean Architecture and SOLID principles
- Analyze time and space complexity where applicable
- Apply patterns in real-world systems and interviews

## Theory

${title} is fundamental to building reliable Go applications. This module progresses from basic concepts to advanced production patterns used at scale.

### Key Concepts

1. **Foundation** — Core syntax, semantics, and idiomatic Go patterns
2. **Advanced Usage** — Edge cases, performance characteristics, and trade-offs
3. **Production Application** — How senior engineers use this in distributed systems

## Visual Diagram

\`\`\`mermaid
flowchart TD
    A[Theory] --> B[Basic Example]
    B --> C[Advanced Example]
    C --> D[Production Implementation]
    D --> E[Interview Discussion]
    E --> F[Real-World Use Case]

    subgraph Learning Path
        B
        C
        D
    end
\`\`\`

## Complexity Analysis

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Basic     | O(1) | O(1)  | See implementation for details |
| Typical   | O(n) | O(n)  | Varies by use case |

## Code Examples

See \`examples/\` for runnable code:

- \`basic.go\` — Minimal introduction
- \`advanced.go\` — Production patterns
- \`main.go\` — Runnable demo

Run:

\`\`\`bash
go run ./${slug}/examples/...
\`\`\`

## Exercises

Complete exercises in \`exercises/\`:

1. **Exercise 1** — Implement the basic pattern from scratch
2. **Exercise 2** — Add error handling and tests
3. **Exercise 3** — Optimize for production workloads

Solutions are in \`exercises/solutions/\`.

## Interview Questions

See \`interview.md\` for curated questions with deep explanations, follow-ups, and production scenarios.

## Production Considerations

See \`production.md\` for:

- Structured logging with context propagation
- Graceful shutdown and health checks
- Configuration management
- Metrics, tracing, and observability
- Security best practices

## Common Mistakes

See \`common-mistakes.md\` — pitfalls that cause bugs, performance issues, and failed interviews.

## Best Practices

See \`best-practices.md\` — idiomatic Go patterns used in production codebases.

## Real-World Use Cases

- **Backend APIs** — REST/gRPC services at scale
- **Distributed Systems** — Microservices, event-driven architectures
- **Infrastructure** — Cloud-native tooling and observability
- **Data Processing** — Pipelines, streaming, and batch workloads

## Related Modules

- Previous: See [150-roadmap](../150-roadmap/README.md) for ordering
- Next: Continue sequentially through the roadmap

## Further Reading

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Blog](https://go.dev/blog/)

---

*Part of [Go Mastery Roadmap](../README.md) — from beginner to staff engineer.*
EOF
}

create_interview_md() {
  local dir="$1" title="$2"
  cat > "$dir/interview.md" << EOF
# ${title} — Interview Questions

## Beginner

### Q1: What is ${title} and why does it matter in Go?

**Answer:** ${title} is a core Go concept used in every production codebase.

**Deep Explanation:** Understanding ${title} demonstrates language fundamentals and informs design decisions in concurrent, distributed systems.

**Follow-up Questions:**
- How would you explain this to a junior developer?
- What Go standard library packages relate to this?

**Production Scenario:** You are building a high-throughput API. How does ${title} affect latency, memory, and correctness?

---

## Intermediate

### Q2: What are common pitfalls with ${title}?

**Answer:** See \`common-mistakes.md\` for detailed pitfalls.

**Deep Explanation:** Most production bugs stem from misunderstanding ownership, concurrency, or error propagation.

**Follow-up Questions:**
- How do you test edge cases?
- What benchmarks would you write?

**Production Scenario:** A service shows elevated p99 latency. How does ${title} factor into your investigation?

---

## Advanced

### Q3: How would you design a system leveraging ${title} at scale?

**Answer:** Apply Clean Architecture layers, context propagation, and observability.

**Deep Explanation:** Staff-level answers cover trade-offs: consistency vs availability, coupling vs cohesion, operational complexity.

**Follow-up Questions:**
- How does this interact with the Go memory model?
- What metrics and alerts would you add?

**Production Scenario:** Design for 10x traffic growth. What changes to your ${title} approach?

---

## Staff Engineer

### Q4: What architectural decisions involve ${title} across a microservices platform?

**Answer:** Cross-cutting concerns include shared libraries, versioning, backward compatibility, and team boundaries.

**Deep Explanation:** Staff engineers evaluate org-wide impact: developer velocity, incident frequency, and total cost of ownership.

**Follow-up Questions:**
- How do you prevent anti-patterns from spreading?
- What would you standardize in a platform team?

**Production Scenario:** Multi-region deployment with strict SLAs — how does ${title} influence your design?
EOF
}

create_production_md() {
  local dir="$1" title="$2"
  cat > "$dir/production.md" << EOF
# ${title} — Production Considerations

## Structured Logging

\`\`\`go
log.Info("operation completed",
    zap.String("module", "${title}"),
    zap.Duration("elapsed", elapsed),
)
\`\`\`

## Context Propagation

Always pass \`context.Context\` as the first parameter. Use it for cancellation, deadlines, and trace IDs.

## Graceful Shutdown

\`\`\`go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
// drain connections, flush buffers
\`\`\`

## Configuration

Use environment variables via \`pkg/config\`. Never hardcode secrets.

## Health Checks

Register dependency checkers with \`pkg/health\` for Kubernetes readiness probes.

## Metrics & Tracing

- Expose Prometheus metrics for latency, throughput, and error rates
- Propagate OpenTelemetry trace context across service boundaries

## Error Handling

Wrap errors with \`pkg/apperrors\`. Return structured errors to callers; log at boundaries.

## Security

- Validate all inputs
- Use parameterized queries for SQL
- Apply least-privilege for service accounts
- Rotate secrets via vault/Kubernetes secrets

## Deployment Checklist

- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Benchmarks within SLA
- [ ] No race conditions (\`go test -race\`)
- [ ] Liveness/readiness probes configured
- [ ] Dashboards and alerts in place
EOF
}

create_common_mistakes() {
  local dir="$1" title="$2"
  cat > "$dir/common-mistakes.md" << EOF
# ${title} — Common Mistakes

## 1. Ignoring Error Returns

Go functions return errors for a reason. Never use \`_ = err\` in production code without explicit justification.

## 2. Missing Context Cancellation

Long-running operations without \`context.Context\` cannot be cancelled during shutdown or timeout.

## 3. Premature Optimization

Optimize based on benchmarks and profiles, not assumptions. Use \`pprof\` and \`benchstat\`.

## 4. Shared Mutable State Without Synchronization

Concurrent access to shared data requires mutexes, channels, or atomic operations. Run \`go test -race\`.

## 5. Interface Pollution

Define interfaces at the consumer, not the producer. Keep interfaces small (1-3 methods).

## 6. Not Using \`defer\` for Cleanup

Always \`defer\` file closes, mutex unlocks, and connection releases to prevent leaks on panic paths.

## 7. Logging Instead of Returning Errors

Log at system boundaries (HTTP handlers, main). Libraries should return errors; callers decide logging.

## 8. Hardcoding Configuration

Use environment variables and config structs. Support 12-factor app principles.
EOF
}

create_best_practices() {
  local dir="$1" title="$2"
  cat > "$dir/best-practices.md" << EOF
# ${title} — Best Practices

## Idiomatic Go

- Accept interfaces, return structs
- Keep functions small and focused (single responsibility)
- Use table-driven tests
- Prefer composition over inheritance

## Clean Architecture

\`\`\`mermaid
flowchart TB
    subgraph External
        HTTP[HTTP Handlers]
        GRPC[gRPC Handlers]
    end
    subgraph Application
        UC[Use Cases]
    end
    subgraph Domain
        ENT[Entities]
        REPO_IF[Repository Interfaces]
    end
    subgraph Infrastructure
        DB[(Database)]
        CACHE[(Cache)]
    end
    HTTP --> UC
    GRPC --> UC
    UC --> ENT
    UC --> REPO_IF
    REPO_IF -.-> DB
    REPO_IF -.-> CACHE
\`\`\`

## SOLID in Go

| Principle | Go Application |
|-----------|----------------|
| Single Responsibility | One package/concern per directory |
| Open/Closed | Extend via interfaces and embedding |
| Liskov Substitution | Interface contracts honored by all impls |
| Interface Segregation | Small, focused interfaces |
| Dependency Inversion | Depend on interfaces in domain layer |

## Testing Strategy

1. **Unit tests** — Pure logic, fast, no I/O
2. **Integration tests** — Real DB/Redis with testcontainers
3. **Benchmarks** — Performance regression detection
4. **Fuzz tests** — Input validation edge cases

## Code Review Checklist

- [ ] Errors handled and wrapped
- [ ] Context propagated
- [ ] No data races
- [ ] Tests cover happy path and edge cases
- [ ] Documentation for exported APIs
EOF
}

pkg_suffix() {
  echo "$1" | sed 's/^[0-9]*-//'
}

create_basic_example() {
  local dir="$1" slug="$2" title="$3"
  local suffix fn_name
  suffix=$(pkg_suffix "$slug")
  fn_name=$(echo "$suffix" | sed 's/-\([a-z]\)/\U\1/g' | sed 's/^\([a-z]\)/\U\1/')
  mkdir -p "$dir/examples"
  cat > "$dir/examples/basic.go" << EOF
// Package examples demonstrates ${title} fundamentals.
package examples

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// BasicDemo runs a minimal ${title} example.
func BasicDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("running basic demo")

	fmt.Printf("=== ${title} — Basic Example ===\\n")
	result := demonstrate${fn_name}()
	fmt.Printf("Result: %v\\n", result)
	return nil
}

func demonstrate${fn_name}() string {
	return "${title} basic demonstration"
}
EOF

  cat > "$dir/examples/advanced.go" << EOF
package examples

import (
	"context"
	"fmt"
	"time"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

// AdvancedDemo shows production patterns for ${title}.
func AdvancedDemo(ctx context.Context) error {
	log := logger.FromContext(ctx)
	start := time.Now()
	defer func() {
		log.Info("advanced demo completed", 
			// elapsed tracked via defer pattern
		)
	}()

	if err := validateContext(ctx); err != nil {
		return err
	}

	fmt.Printf("=== ${title} — Advanced Example ===\\n")
	fmt.Printf("Elapsed: %v\\n", time.Since(start))
	return nil
}

func validateContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "context cancelled")
	default:
		return nil
	}
}
EOF

  cat > "$dir/examples/main.go" << EOF
package main

import (
	"context"
	"os"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}/examples"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/logger"
)

func main() {
	log := logger.Must("${slug}")
	ctx := logger.WithContext(context.Background(), log)

	if err := examples.BasicDemo(ctx); err != nil {
		log.Error("basic demo failed")
		os.Exit(1)
	}
	if err := examples.AdvancedDemo(ctx); err != nil {
		log.Error("advanced demo failed")
		os.Exit(1)
	}
}
EOF
}

create_exercises() {
  local dir="$1" title="$2"
  mkdir -p "$dir/exercises/solutions"
  cat > "$dir/exercises/README.md" << EOF
# ${title} — Exercises

## Exercise 1: Fundamentals

Implement the core concept from scratch without looking at solutions.

**Requirements:**
- Write idiomatic Go code
- Include error handling
- Add at least 3 unit tests

## Exercise 2: Production Patterns

Extend your implementation with:
- Context support
- Structured logging
- Graceful error wrapping

## Exercise 3: Performance

- Write benchmarks
- Compare at least 2 approaches
- Document complexity analysis

## Exercise 4: Interview Challenge

Solve the problem in \`interview-challenge.go\` within 30 minutes.

Check \`solutions/\` after attempting all exercises.
EOF

  cat > "$dir/exercises/exercise_test.go" << 'EOF'
package exercises_test

import "testing"

func TestExercise1Placeholder(t *testing.T) {
	t.Skip("Implement Exercise 1, then remove this skip")
}
EOF
}

create_ds_files() {
  local dir="$1" slug="$2" title="$3"
  local pkg
  pkg=$(pkg_suffix "$slug" | tr '-' '_')

  cat > "$dir/implementation.go" << EOF
// Package ${pkg} implements ${title} from scratch for learning purposes.
package ${pkg}

import "errors"

var ErrEmpty = errors.New("${pkg}: empty structure")

// DataStructure is a placeholder — see module-specific implementations.
type DataStructure struct {
	data []int
}

// New creates a new ${title}.
func New() *DataStructure {
	return &DataStructure{data: make([]int, 0)}
}

// Size returns the number of elements.
func (ds *DataStructure) Size() int {
	return len(ds.data)
}

// IsEmpty reports whether the structure has no elements.
func (ds *DataStructure) IsEmpty() bool {
	return len(ds.data) == 0
}
EOF

  cat > "$dir/implementation_test.go" << EOF
package ${pkg}_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}"
)

func TestNew(t *testing.T) {
	ds := ${pkg}.New()
	if !ds.IsEmpty() {
		t.Fatal("expected empty structure")
	}
	if ds.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ds.Size())
	}
}
EOF

  cat > "$dir/benchmarks_test.go" << EOF
package ${pkg}_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ${pkg}.New()
	}
}
EOF
}

for entry in "${MODULES[@]}"; do
  IFS='|' read -r slug title category level <<< "$entry"
  dir="$ROOT/$slug"
  mkdir -p "$dir/examples" "$dir/exercises/solutions" "$dir/diagrams"

  create_readme "$dir" "$slug" "$title" "$category" "$level"
  create_interview_md "$dir" "$title"
  create_production_md "$dir" "$title"
  create_common_mistakes "$dir" "$title"
  create_best_practices "$dir" "$title"
  create_basic_example "$dir" "$slug" "$title"
  create_exercises "$dir" "$title"

  # Mermaid diagram file
  cat > "$dir/diagrams/overview.mmd" << EOF
flowchart LR
    Input[Input] --> Process[${title}]
    Process --> Output[Output]
EOF

  if is_ds_module "$slug"; then
    create_ds_files "$dir" "$slug" "$title"
  fi

  if is_algo_module "$slug"; then
    create_ds_files "$dir" "$slug" "$title"
    cat > "$dir/complexity.md" << EOF
# ${title} — Complexity Analysis

| Algorithm/Operation | Best | Average | Worst | Space |
|---------------------|------|---------|-------|-------|
| Primary operation   | O(1) | O(n log n) | O(n²) | O(n) |

See \`implementation.go\` for algorithm-specific analysis.
EOF
  fi

  echo "Created: $slug"
done

echo ""
echo "Generated ${#MODULES[@]} modules successfully."
