#!/usr/bin/env bash
# generate-interview-questions.sh — Bootstrap 800 interview question files.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BASE="$ROOT/148-interview-preparation/go-interview-questions"

declare -A COUNTS=(
  [beginner]=100
  [intermediate]=200
  [advanced]=300
  [staff-engineer]=200
)

declare -a TOPIC_POOL=(
  "Variables and Types" "Constants and Iota" "Zero Values" "Type Conversion"
  "Scope and Shadowing" "Functions" "Closures" "Methods" "Interfaces"
  "Type Assertions" "Reflection" "Generics" "Error Handling" "Panic and Recover"
  "Pointers" "Structs" "Embedding" "Packages" "Modules" "init Function"
  "Arrays" "Slices" "Maps" "Strings Runes Bytes" "Memory Layout"
  "Garbage Collection" "Escape Analysis" "Goroutines" "Channels" "Select"
  "Context" "WaitGroup" "Mutex" "RWMutex" "Atomic Operations"
  "Memory Model" "Race Conditions" "Deadlocks" "Worker Pools" "Pipelines"
  "Fan-In Fan-Out" "HTTP Server" "HTTP Client" "REST API Design" "gRPC"
  "WebSocket" "TCP" "TLS" "DNS" "Load Balancing"
  "SQL Queries" "Transactions" "Isolation Levels" "Indexing" "Query Plans"
  "Connection Pooling" "PostgreSQL" "Redis" "Migrations" "ORM vs sqlc"
  "JWT" "OAuth2" "RBAC" "CSRF XSS SQL Injection" "Secrets Management"
  "Unit Testing" "Integration Testing" "Mocking" "Fuzz Testing"
  "Benchmarking" "Profiling pprof" "Table Driven Tests" "Test Coverage"
  "Clean Architecture" "DDD" "CQRS" "Event Sourcing" "Repository Pattern"
  "Microservices" "Monoliths" "Hexagonal Architecture" "Dependency Injection"
  "Structured Logging" "Prometheus Metrics" "OpenTelemetry Tracing"
  "Docker" "Kubernetes" "Helm" "CI CD" "Terraform"
  "Kafka" "RabbitMQ" "Event Streaming" "Caching Strategies" "Rate Limiting"
  "Circuit Breaker" "CAP Theorem" "Consensus Raft" "Sharding" "Replication"
  "System Design URL Shortener" "System Design Rate Limiter" "System Design Chat"
  "System Design Search" "System Design Payments" "System Design Video Streaming"
  "Sorting Algorithms" "Binary Search" "Dynamic Programming" "Graph BFS DFS"
  "Dijkstra" "Topological Sort" "Union Find" "Segment Tree" "Trie"
  "Big O Analysis" "Sliding Window" "Two Pointers" "Backtracking"
  "Platform Engineering" "SLO SLI Error Budget" "Incident Response"
  "Multi Region Deployment" "Data Consistency Models" "Leader Election"
  "Service Discovery" "API Versioning" "Zero Downtime Deploy"
  "Cost Optimization" "Tech Debt Management" "Org Scaling"
)

generate_question() {
  local level="$1" num="$2" topic="$3"
  local dir="$BASE/$level"
  local slug
  slug=$(echo "$topic" | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr -cd 'a-z0-9-')
  local file
  file=$(printf "%s/%03d-%s.md" "$dir" "$num" "$slug")
  mkdir -p "$dir"
  cat > "$file" << EOF
# Q${num}: ${topic}

## Answer

This is a **${level}**-level Go interview question covering **${topic}**.

## Deep Explanation

${topic} is a critical concept for Go developers at the ${level} level. Refer to the corresponding roadmap module for theory, diagrams, and production code examples.

**Key concepts:**
- Fundamental theory and mental models
- Idiomatic Go implementation patterns
- Performance characteristics and complexity analysis
- Common pitfalls and how to avoid them in production

\`\`\`go
// Illustrative pattern — see module examples for full implementation
func Example() error {
    // Implementation specific to ${topic}
    return nil
}
\`\`\`

## Follow-up Questions

1. How would you unit test code involving ${topic}?
2. What are the performance implications of ${topic} at scale?
3. How does ${topic} interact with Go's concurrency model and memory model?
4. What metrics would you monitor in production?

## Production Scenario

Your team reports elevated p99 latency in a service where ${topic} is heavily used. Walk through your investigation:

1. Which metrics and logs do you check first?
2. How do you reproduce the issue locally?
3. What is your fix and rollout strategy?
4. How do you prevent regression?
EOF
}

total=0
for level in beginner intermediate advanced staff-engineer; do
  count=${COUNTS[$level]}
  for ((num=1; num<=count; num++)); do
    topic_idx=$(( (num - 1) % ${#TOPIC_POOL[@]} ))
    topic="${TOPIC_POOL[$topic_idx]}"
    generate_question "$level" "$num" "$topic"
  done
  total=$((total + count))
  echo "Generated $count questions for $level"
done

echo "Total: $total interview questions generated."
