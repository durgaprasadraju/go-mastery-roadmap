// Command generate-exercise-solutions writes topic-specific exercise solutions per module.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type module struct {
	slug, title, category string
}

var modules = []module{
	{"01-fundamentals", "Go Fundamentals", "fundamentals"},
	{"02-types-system", "Go Type System", "fundamentals"},
	{"03-control-flow", "Control Flow", "fundamentals"},
	{"04-functions", "Functions & Closures", "fundamentals"},
	{"05-packages-modules", "Packages & Modules", "fundamentals"},
	{"06-pointers", "Pointers", "fundamentals"},
	{"07-structs", "Structs", "fundamentals"},
	{"08-methods", "Methods", "fundamentals"},
	{"09-interfaces", "Interfaces", "fundamentals"},
	{"10-generics", "Generics", "fundamentals"},
	{"11-error-handling", "Error Handling", "fundamentals"},
	{"12-memory-management", "Memory Management", "fundamentals"},
	{"13-arrays", "Arrays", "data-structures"},
	{"14-slices", "Slices", "data-structures"},
	{"15-maps", "Maps", "data-structures"},
	{"16-strings-runes-bytes", "Strings, Runes & Bytes", "data-structures"},
	{"17-recursion", "Recursion", "data-structures"},
	{"18-linked-lists", "Linked Lists", "data-structures"},
	{"19-stacks", "Stacks", "data-structures"},
	{"20-queues", "Queues", "data-structures"},
	{"21-trees", "Trees", "data-structures"},
	{"22-binary-search-tree", "Binary Search Trees", "data-structures"},
	{"23-heaps", "Heaps", "data-structures"},
	{"24-priority-queue", "Priority Queues", "data-structures"},
	{"25-hash-tables", "Hash Tables", "data-structures"},
	{"26-trie", "Tries", "data-structures"},
	{"27-graphs", "Graphs", "data-structures"},
	{"28-union-find", "Union-Find", "data-structures"},
	{"29-segment-tree", "Segment Trees", "data-structures"},
	{"30-fenwick-tree", "Fenwick Trees", "data-structures"},
	{"31-sorting", "Sorting Algorithms", "algorithms"},
	{"32-searching", "Searching Algorithms", "algorithms"},
	{"33-sliding-window", "Sliding Window", "algorithms"},
	{"34-two-pointers", "Two Pointers", "algorithms"},
	{"35-fast-slow-pointers", "Fast & Slow Pointers", "algorithms"},
	{"36-monotonic-stack", "Monotonic Stack", "algorithms"},
	{"37-monotonic-queue", "Monotonic Queue", "algorithms"},
	{"38-backtracking", "Backtracking", "algorithms"},
	{"39-divide-conquer", "Divide & Conquer", "algorithms"},
	{"40-greedy", "Greedy Algorithms", "algorithms"},
	{"41-dynamic-programming", "Dynamic Programming", "algorithms"},
	{"42-bit-manipulation", "Bit Manipulation", "algorithms"},
	{"43-math-algorithms", "Math Algorithms", "algorithms"},
	{"44-string-algorithms", "String Algorithms", "algorithms"},
	{"45-graph-algorithms", "Graph Algorithms", "algorithms"},
	{"46-tree-algorithms", "Tree Algorithms", "algorithms"},
	{"47-big-o-analysis", "Big-O Analysis", "complexity"},
	{"48-complexity-analysis", "Complexity Analysis", "complexity"},
	{"49-goroutines", "Goroutines", "concurrency"},
	{"50-channels", "Channels", "concurrency"},
	{"51-worker-pools", "Worker Pools", "concurrency"},
	{"52-pipelines", "Pipelines", "concurrency"},
	{"53-fan-in-fan-out", "Fan-In / Fan-Out", "concurrency"},
	{"54-context", "Context", "concurrency"},
	{"55-select", "Select Statement", "concurrency"},
	{"56-sync-package", "sync Package", "concurrency"},
	{"57-mutex", "Mutex", "concurrency"},
	{"58-rwmutex", "RWMutex", "concurrency"},
	{"59-cond", "sync.Cond", "concurrency"},
	{"60-atomic", "Atomic Operations", "concurrency"},
	{"61-memory-model", "Go Memory Model", "concurrency"},
	{"62-race-conditions", "Race Conditions", "concurrency"},
	{"63-deadlocks", "Deadlocks", "concurrency"},
	{"64-concurrency-patterns", "Concurrency Patterns", "concurrency"},
	{"65-file-handling", "File Handling", "io"},
	{"66-json", "JSON", "io"},
	{"67-yaml", "YAML", "io"},
	{"68-csv", "CSV", "io"},
	{"69-xml", "XML", "io"},
	{"70-compression", "Compression", "io"},
	{"71-networking-fundamentals", "Networking Fundamentals", "networking"},
	{"72-tcp", "TCP", "networking"},
	{"73-udp", "UDP", "networking"},
	{"74-http", "HTTP", "networking"},
	{"75-http2", "HTTP/2", "networking"},
	{"76-http3", "HTTP/3", "networking"},
	{"77-rest-api", "REST API", "networking"},
	{"78-grpc", "gRPC", "networking"},
	{"79-websocket", "WebSocket", "networking"},
	{"80-server-sent-events", "Server-Sent Events", "networking"},
	{"81-dns", "DNS", "networking"},
	{"82-load-balancing", "Load Balancing", "networking"},
	{"83-database-sql", "SQL Fundamentals", "database"},
	{"84-postgresql", "PostgreSQL", "database"},
	{"85-mysql", "MySQL", "database"},
	{"86-sqlite", "SQLite", "database"},
	{"87-transactions", "Transactions", "database"},
	{"88-indexing", "Indexing", "database"},
	{"89-query-optimization", "Query Optimization", "database"},
	{"90-connection-pooling", "Connection Pooling", "database"},
	{"91-orm-gorm", "GORM ORM", "database"},
	{"92-sqlc", "sqlc", "database"},
	{"93-migrations", "Database Migrations", "database"},
	{"94-authentication", "Authentication", "security"},
	{"95-authorization", "Authorization", "security"},
	{"96-jwt", "JWT", "security"},
	{"97-oauth2", "OAuth2", "security"},
	{"98-session-management", "Session Management", "security"},
	{"99-security", "Security Fundamentals", "security"},
	{"100-encryption", "Encryption", "security"},
	{"101-hashing", "Hashing", "security"},
	{"102-secrets-management", "Secrets Management", "security"},
	{"103-testing", "Testing Fundamentals", "testing"},
	{"104-unit-testing", "Unit Testing", "testing"},
	{"105-integration-testing", "Integration Testing", "testing"},
	{"106-e2e-testing", "E2E Testing", "testing"},
	{"107-mocking", "Mocking", "testing"},
	{"108-fuzz-testing", "Fuzz Testing", "testing"},
	{"109-benchmarking", "Benchmarking", "testing"},
	{"110-profiling", "Profiling", "testing"},
	{"111-clean-architecture", "Clean Architecture", "architecture"},
	{"112-ddd", "Domain-Driven Design", "architecture"},
	{"113-cqrs", "CQRS", "architecture"},
	{"114-event-driven", "Event-Driven Architecture", "architecture"},
	{"115-microservices", "Microservices", "architecture"},
	{"116-monoliths", "Monoliths", "architecture"},
	{"117-hexagonal", "Hexagonal Architecture", "architecture"},
	{"118-repository-pattern", "Repository Pattern", "architecture"},
	{"119-logging", "Logging", "observability"},
	{"120-observability", "Observability", "observability"},
	{"121-prometheus", "Prometheus", "observability"},
	{"122-grafana", "Grafana", "observability"},
	{"123-opentelemetry", "OpenTelemetry", "observability"},
	{"124-tracing", "Distributed Tracing", "observability"},
	{"125-docker", "Docker", "cloud-native"},
	{"126-kubernetes", "Kubernetes", "cloud-native"},
	{"127-helm", "Helm", "cloud-native"},
	{"128-ci-cd", "CI/CD", "cloud-native"},
	{"129-github-actions", "GitHub Actions", "cloud-native"},
	{"130-terraform", "Terraform", "cloud-native"},
	{"131-redis", "Redis", "messaging"},
	{"132-kafka", "Apache Kafka", "messaging"},
	{"133-rabbitmq", "RabbitMQ", "messaging"},
	{"134-nats", "NATS", "messaging"},
	{"135-event-streaming", "Event Streaming", "messaging"},
	{"136-design-patterns", "Design Patterns", "patterns"},
	{"137-system-design", "System Design", "system-design"},
	{"138-distributed-systems", "Distributed Systems", "system-design"},
	{"139-consensus", "Consensus Algorithms", "system-design"},
	{"140-raft", "Raft Consensus", "system-design"},
	{"141-cap-theorem", "CAP Theorem", "system-design"},
	{"142-sharding", "Sharding", "system-design"},
	{"143-replication", "Replication", "system-design"},
	{"144-caching", "Caching Strategies", "system-design"},
	{"145-rate-limiting", "Rate Limiting", "system-design"},
	{"146-circuit-breaker", "Circuit Breaker", "system-design"},
	{"147-production-projects", "Production Projects", "projects"},
	{"148-interview-preparation", "Interview Preparation", "interview"},
	{"149-cheat-sheets", "Cheat Sheets", "reference"},
	{"150-roadmap", "Learning Roadmap", "reference"},
}

func main() {
	root, _ := os.Getwd()
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	for _, m := range modules {
		if m.slug == "01-fundamentals" {
			fmt.Println("skip (custom):", m.slug)
			continue
		}
		if err := writeModule(root, m); err != nil {
			fmt.Fprintf(os.Stderr, "error %s: %v\n", m.slug, err)
			os.Exit(1)
		}
		fmt.Println("generated:", m.slug)
	}
}

func writeModule(root string, m module) error {
	dir := filepath.Join(root, m.slug, "exercises", "solutions")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	spec := buildSpec(m)
	files := map[string]string{
		"README.md":                         readme(m, spec),
		"exercise1.go":                      spec.exercise1,
		"exercise1_test.go":                 spec.exercise1Test,
		"exercise2.go":                      exercise2(m, spec),
		"exercise2_test.go":                 exercise2Test(m),
		"exercise3_bench_test.go":           exercise3Bench(m, spec),
		"interview-challenge-solution.go":   spec.interviewSolution,
		"interview-challenge-solution_test.go": spec.interviewTest,
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			return err
		}
	}
	challenge := filepath.Join(root, m.slug, "exercises", "interview-challenge.go")
	if err := os.WriteFile(challenge, []byte(interviewChallengeStub(m, spec)), 0o644); err != nil {
		return err
	}
	exercisesReadme := filepath.Join(root, m.slug, "exercises", "README.md")
	if m.slug == "01-fundamentals" {
		return nil // keep custom fundamentals exercises README
	}
	return os.WriteFile(exercisesReadme, []byte(exercisesReadmeContent(m, spec)), 0o644)
}

type topicSpec struct {
	summary            string
	exercise1          string
	exercise1Test      string
	interviewProblem   string
	interviewSolution  string
	interviewTest      string
	benchCall          string
	testCall           string
}

func pkgSuffix(slug string) string {
	return strings.TrimPrefix(slug, strings.Split(slug, "-")[0]+"-")
}

func errPrefix(slug string) string {
	return strings.ReplaceAll(pkgSuffix(slug), "-", "")
}

func modImport(slug string) string {
	return fmt.Sprintf("github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions", slug)
}

func buildSpec(m module) topicSpec {
	var spec topicSpec
	if s, ok := overrides[m.slug]; ok {
		spec = s
	} else {
		switch m.category {
		case "fundamentals":
			spec = fundamentalsSpec(m)
		case "data-structures":
			spec = dataStructureSpec(m)
		case "algorithms":
			spec = algorithmSpec(m)
		case "complexity":
			spec = complexitySpec(m)
		case "concurrency":
			spec = concurrencySpec(m)
		case "io":
			spec = ioSpec(m)
		case "networking":
			spec = networkingSpec(m)
		case "database":
			spec = databaseSpec(m)
		case "security":
			spec = securitySpec(m)
		case "testing":
			spec = testingSpec(m)
		case "architecture":
			spec = architectureSpec(m)
		case "observability":
			spec = observabilitySpec(m)
		case "cloud-native":
			spec = cloudSpec(m)
		case "messaging":
			spec = messagingSpec(m)
		case "patterns", "system-design":
			spec = systemDesignSpec(m)
		default:
			spec = metaSpec(m)
		}
	}
	return finalizeSpec(m, spec)
}
