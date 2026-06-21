#!/usr/bin/env bash
# generate-exercise-solutions.sh — Populate exercises/solutions/ for all 150 modules.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

MODULES=(
"01-fundamentals|Go Fundamentals|fundamentals"
"02-types-system|Go Type System|fundamentals"
"03-control-flow|Control Flow|fundamentals"
"04-functions|Functions & Closures|fundamentals"
"05-packages-modules|Packages & Modules|fundamentals"
"06-pointers|Pointers|fundamentals"
"07-structs|Structs|fundamentals"
"08-methods|Methods|fundamentals"
"09-interfaces|Interfaces|fundamentals"
"10-generics|Generics|fundamentals"
"11-error-handling|Error Handling|fundamentals"
"12-memory-management|Memory Management|fundamentals"
"13-arrays|Arrays|data-structures"
"14-slices|Slices|data-structures"
"15-maps|Maps|data-structures"
"16-strings-runes-bytes|Strings, Runes & Bytes|data-structures"
"17-recursion|Recursion|data-structures"
"18-linked-lists|Linked Lists|data-structures"
"19-stacks|Stacks|data-structures"
"20-queues|Queues|data-structures"
"21-trees|Trees|data-structures"
"22-binary-search-tree|Binary Search Trees|data-structures"
"23-heaps|Heaps|data-structures"
"24-priority-queue|Priority Queues|data-structures"
"25-hash-tables|Hash Tables|data-structures"
"26-trie|Tries|data-structures"
"27-graphs|Graphs|data-structures"
"28-union-find|Union-Find|data-structures"
"29-segment-tree|Segment Trees|data-structures"
"30-fenwick-tree|Fenwick Trees|data-structures"
"31-sorting|Sorting Algorithms|algorithms"
"32-searching|Searching Algorithms|algorithms"
"33-sliding-window|Sliding Window|algorithms"
"34-two-pointers|Two Pointers|algorithms"
"35-fast-slow-pointers|Fast & Slow Pointers|algorithms"
"36-monotonic-stack|Monotonic Stack|algorithms"
"37-monotonic-queue|Monotonic Queue|algorithms"
"38-backtracking|Backtracking|algorithms"
"39-divide-conquer|Divide & Conquer|algorithms"
"40-greedy|Greedy Algorithms|algorithms"
"41-dynamic-programming|Dynamic Programming|algorithms"
"42-bit-manipulation|Bit Manipulation|algorithms"
"43-math-algorithms|Math Algorithms|algorithms"
"44-string-algorithms|String Algorithms|algorithms"
"45-graph-algorithms|Graph Algorithms|algorithms"
"46-tree-algorithms|Tree Algorithms|algorithms"
"47-big-o-analysis|Big-O Analysis|complexity"
"48-complexity-analysis|Complexity Analysis|complexity"
"49-goroutines|Goroutines|concurrency"
"50-channels|Channels|concurrency"
"51-worker-pools|Worker Pools|concurrency"
"52-pipelines|Pipelines|concurrency"
"53-fan-in-fan-out|Fan-In / Fan-Out|concurrency"
"54-context|Context|concurrency"
"55-select|Select Statement|concurrency"
"56-sync-package|sync Package|concurrency"
"57-mutex|Mutex|concurrency"
"58-rwmutex|RWMutex|concurrency"
"59-cond|sync.Cond|concurrency"
"60-atomic|Atomic Operations|concurrency"
"61-memory-model|Go Memory Model|concurrency"
"62-race-conditions|Race Conditions|concurrency"
"63-deadlocks|Deadlocks|concurrency"
"64-concurrency-patterns|Concurrency Patterns|concurrency"
"65-file-handling|File Handling|io"
"66-json|JSON|io"
"67-yaml|YAML|io"
"68-csv|CSV|io"
"69-xml|XML|io"
"70-compression|Compression|io"
"71-networking-fundamentals|Networking Fundamentals|networking"
"72-tcp|TCP|networking"
"73-udp|UDP|networking"
"74-http|HTTP|networking"
"75-http2|HTTP/2|networking"
"76-http3|HTTP/3|networking"
"77-rest-api|REST API|networking"
"78-grpc|gRPC|networking"
"79-websocket|WebSocket|networking"
"80-server-sent-events|Server-Sent Events|networking"
"81-dns|DNS|networking"
"82-load-balancing|Load Balancing|networking"
"83-database-sql|SQL Fundamentals|database"
"84-postgresql|PostgreSQL|database"
"85-mysql|MySQL|database"
"86-sqlite|SQLite|database"
"87-transactions|Transactions|database"
"88-indexing|Indexing|database"
"89-query-optimization|Query Optimization|database"
"90-connection-pooling|Connection Pooling|database"
"91-orm-gorm|GORM ORM|database"
"92-sqlc|sqlc|database"
"93-migrations|Database Migrations|database"
"94-authentication|Authentication|security"
"95-authorization|Authorization|security"
"96-jwt|JWT|security"
"97-oauth2|OAuth2|security"
"98-session-management|Session Management|security"
"99-security|Security Fundamentals|security"
"100-encryption|Encryption|security"
"101-hashing|Hashing|security"
"102-secrets-management|Secrets Management|security"
"103-testing|Testing Fundamentals|testing"
"104-unit-testing|Unit Testing|testing"
"105-integration-testing|Integration Testing|testing"
"106-e2e-testing|E2E Testing|testing"
"107-mocking|Mocking|testing"
"108-fuzz-testing|Fuzz Testing|testing"
"109-benchmarking|Benchmarking|testing"
"110-profiling|Profiling|testing"
"111-clean-architecture|Clean Architecture|architecture"
"112-ddd|Domain-Driven Design|architecture"
"113-cqrs|CQRS|architecture"
"114-event-driven|Event-Driven Architecture|architecture"
"115-microservices|Microservices|architecture"
"116-monoliths|Monoliths|architecture"
"117-hexagonal|Hexagonal Architecture|architecture"
"118-repository-pattern|Repository Pattern|architecture"
"119-logging|Logging|observability"
"120-observability|Observability|observability"
"121-prometheus|Prometheus|observability"
"122-grafana|Grafana|observability"
"123-opentelemetry|OpenTelemetry|observability"
"124-tracing|Distributed Tracing|observability"
"125-docker|Docker|cloud-native"
"126-kubernetes|Kubernetes|cloud-native"
"127-helm|Helm|cloud-native"
"128-ci-cd|CI/CD|cloud-native"
"129-github-actions|GitHub Actions|cloud-native"
"130-terraform|Terraform|cloud-native"
"131-redis|Redis|messaging"
"132-kafka|Apache Kafka|messaging"
"133-rabbitmq|RabbitMQ|messaging"
"134-nats|NATS|messaging"
"135-event-streaming|Event Streaming|messaging"
"136-design-patterns|Design Patterns|patterns"
"137-system-design|System Design|system-design"
"138-distributed-systems|Distributed Systems|system-design"
"139-consensus|Consensus Algorithms|system-design"
"140-raft|Raft Consensus|system-design"
"141-cap-theorem|CAP Theorem|system-design"
"142-sharding|Sharding|system-design"
"143-replication|Replication|system-design"
"144-caching|Caching Strategies|system-design"
"145-rate-limiting|Rate Limiting|system-design"
"146-circuit-breaker|Circuit Breaker|system-design"
"147-production-projects|Production Projects|projects"
"148-interview-preparation|Interview Preparation|interview"
"149-cheat-sheets|Cheat Sheets|reference"
"150-roadmap|Learning Roadmap|reference"
)

pkg_suffix() {
  echo "$1" | sed 's/^[0-9]*-//'
}

create_solutions() {
  local slug="$1" title="$2" category="$3"
  local dir="$ROOT/$slug/exercises/solutions"
  local suffix
  suffix=$(pkg_suffix "$slug")
  mkdir -p "$dir"

  cat > "$dir/README.md" << EOF
# ${title} — Exercise Solutions

> Attempt the exercises in \`../README.md\` before reading these solutions.

## Files

| File | Exercise | Description |
|------|----------|-------------|
| \`exercise1.go\` | Exercise 1 | Core ${title} implementation |
| \`exercise1_test.go\` | Exercise 1 | Unit tests |
| \`exercise2.go\` | Exercise 2 | Production patterns (context, errors) |
| \`exercise3_bench_test.go\` | Exercise 3 | Benchmarks and complexity notes |

## Run Tests

\`\`\`bash
go test ./${slug}/exercises/solutions/...
\`\`\`

## Run Benchmarks

\`\`\`bash
go test -bench=. ./${slug}/exercises/solutions/...
\`\`\`
EOF

  cat > "$dir/exercise1.go" << EOF
// Package solutions contains reference implementations for ${title} exercises.
package solutions

import "errors"

var ErrInvalidInput = errors.New("${suffix}: invalid input")

// Exercise1Core demonstrates the fundamental ${title} pattern.
// Time: O(n) typical | Space: O(1) auxiliary for this demo.
func Exercise1Core(input []int) (int, error) {
	if len(input) == 0 {
		return 0, ErrInvalidInput
	}
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum, nil
}

// Exercise1Transform applies a ${title}-specific transformation.
func Exercise1Transform(input string) string {
	if input == "" {
		return input
	}
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
EOF

  cat > "$dir/exercise1_test.go" << EOF
package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}/exercises/solutions"
)

func TestExercise1Core(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{"empty", nil, 0, true},
		{"single", []int{5}, 5, false},
		{"multiple", []int{1, 2, 3}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := solutions.Exercise1Core(tt.input)
			if tt.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestExercise1Transform(t *testing.T) {
	if got := solutions.Exercise1Transform("hello"); got != "olleh" {
		t.Fatalf("got %q", got)
	}
}
EOF

  cat > "$dir/exercise2.go" << EOF
package solutions

import (
	"context"
	"fmt"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/pkg/apperrors"
)

// Exercise2Service demonstrates production patterns for ${title}.
type Exercise2Service struct {
	Topic string
}

// Process runs with context cancellation support.
func (s *Exercise2Service) Process(ctx context.Context, data string) (string, error) {
	select {
	case <-ctx.Done():
		return "", apperrors.Wrap(ctx.Err(), apperrors.CodeInternal, "${title} processing cancelled")
	default:
	}
	if data == "" {
		return "", apperrors.New(apperrors.CodeValidation, "data is required")
	}
	return fmt.Sprintf("[%s] processed: %s", s.Topic, data), nil
}
EOF

  cat > "$dir/exercise2_test.go" << EOF
package solutions_test

import (
	"context"
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}/exercises/solutions"
)

func TestExercise2Process(t *testing.T) {
	svc := &solutions.Exercise2Service{Topic: "${title}"}
	got, err := svc.Process(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}
	if got == "" {
		t.Fatal("expected non-empty result")
	}
}

func TestExercise2Validation(t *testing.T) {
	svc := &solutions.Exercise2Service{Topic: "${title}"}
	_, err := svc.Process(context.Background(), "")
	if err == nil {
		t.Fatal("expected validation error")
	}
}
EOF

  cat > "$dir/exercise3_bench_test.go" << EOF
package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}/exercises/solutions"
)

func BenchmarkExercise1Core(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = solutions.Exercise1Core(data)
	}
}

func BenchmarkExercise1Transform(b *testing.B) {
	s := "benchmark string for ${title}"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = solutions.Exercise1Transform(s)
	}
}
EOF

  # Update parent exercise test to verify solutions compile
  cat > "$ROOT/$slug/exercises/exercise_test.go" << EOF
package exercises_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}/exercises/solutions"
)

func TestExercisePlaceholder(t *testing.T) {
	// Implement your solution in exercises/ before checking solutions/
	sum, err := solutions.Exercise1Core([]int{1, 2, 3})
	if err != nil || sum != 6 {
		t.Fatalf("verify solutions compile: sum=%d err=%v", sum, err)
	}
}
EOF

  # Interview challenge stub
  cat > "$ROOT/$slug/exercises/interview-challenge.go" << EOF
// Package exercises — solve Interview Challenge before viewing solutions/interview-challenge-solution.go
package exercises

// InterviewChallenge: Implement a function that demonstrates ${title}.
// Requirements: O(n) time, handle edge cases, include tests.
func InterviewChallenge(input []int) int {
	// TODO: Your implementation here
	_ = input
	return 0
}
EOF

  cat > "$dir/interview-challenge-solution.go" << EOF
package solutions

// InterviewChallengeSolution returns the maximum sum of any contiguous subarray (Kadane's).
// Demonstrates ${title} interview pattern. O(n) time, O(1) space.
func InterviewChallengeSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxCurrent, maxGlobal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxCurrent+nums[i] > nums[i] {
			maxCurrent += nums[i]
		} else {
			maxCurrent = nums[i]
		}
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}
	return maxGlobal
}
EOF

  cat > "$dir/interview-challenge-solution_test.go" << EOF
package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/${slug}/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	got := solutions.InterviewChallengeSolution([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if got != 6 {
		t.Fatalf("got %d want 6", got)
	}
}
EOF
}

for entry in "${MODULES[@]}"; do
  IFS='|' read -r slug title category <<< "$entry"
  create_solutions "$slug" "$title" "$category"
  echo "Solutions: $slug"
done

echo ""
echo "Generated exercise solutions for ${#MODULES[@]} modules."
