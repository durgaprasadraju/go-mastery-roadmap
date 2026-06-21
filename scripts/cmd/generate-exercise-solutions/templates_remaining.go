package main

import "fmt"

func concurrencySpec(m module) topicSpec {
	cases := map[string]struct{ body, test, problem, interview, iTest, bench string }{
		"49-goroutines": {
			goroutineBody(), testGoroutine(m), "Parallel sum with goroutines + WaitGroup",
			pingPongInterviewBody(), testPingPongInterview(m), `solutions.ParallelSum([]int{1,2,3,4}, 2)`,
		},
		"50-channels": {
			channelBody(), testChannel(m), "Producer-consumer with buffered channel",
			producerConsumerBody(), testProducerConsumer(m), `PingPong(3)`,
		},
		"51-worker-pools": {
			workerPoolBody(), testWorkerPool(m), "Worker pool processing jobs",
			workerPoolBody(), testWorkerPool(m), `RunPool(10, 4)`,
		},
		"54-context": {
			contextBody(), testContext(m), "Propagate context cancellation",
			contextBody(), testContext(m), `WithTimeoutOp(context.Background(), 1)`,
		},
		"57-mutex": {
			mutexBody(), testMutex(m), "Thread-safe counter with Mutex",
			mutexBody(), testMutex(m), `SafeCounter.Inc()`,
		},
		"60-atomic": {
			atomicBody(), testAtomic(m), "Lock-free counter with atomic.Int64",
			atomicBody(), testAtomic(m), `AtomicCounter.Add(1)`,
		},
	}
	if c, ok := cases[m.slug]; ok {
		return topicSpec{summary: fmt.Sprintf("**%s** concurrency patterns in Go.", m.title), exercise1: c.body, exercise1Test: c.test, interviewProblem: c.problem, interviewSolution: c.interview, interviewTest: c.iTest, benchCall: c.bench}
	}
	interview := pingPongInterviewBody()
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — goroutines, channels, sync primitives.", m.title),
		exercise1:         goroutineBody(),
		exercise1Test:     testGoroutine(m),
		interviewProblem:  "Coordinate concurrent work safely",
		interviewSolution: interview,
		interviewTest:     testPingPongInterview(m),
		benchCall:         `solutions.ParallelSum([]int{1,2,3}, 2)`,
	}
}

func pingPongInterviewBody() string {
	return `package solutions

// InterviewChallengeSolution runs n synchronous ping-pong exchanges.
func InterviewChallengeSolution(n int) int {
	ch := make(chan int, 1)
	ch <- 0
	count := 0
	for i := 0; i < n; i++ {
		<-ch
		count++
		ch <- count
	}
	return <-ch
}`
}

func networkingSpec(m module) topicSpec {
	body := `package solutions

import "strings"

// ParseHostPort splits host:port (simplified).
func ParseHostPort(addr string) (host, port string, ok bool) {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

// BuildURL builds http URL from parts.
func BuildURL(host, path string) string {
	return "http://" + host + path
}`
	test := fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestNetworking(t *testing.T) {
	h, p, ok := solutions.ParseHostPort("localhost:8080")
	if !ok || h != "localhost" || p != "8080" {
		t.Fatalf("host=%%s port=%%s ok=%%v", h, p, ok)
	}
	if solutions.BuildURL("api.local", "/v1") != "http://api.local/v1" {
		t.Fatal("build url failed")
	}
}
`, m.slug)
	interview := restRouteMatchBody()
	interviewTest := testRestRouteMatch(m)
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — protocols, clients, servers.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Design HTTP handler with middleware chain",
		interviewSolution: interview,
		interviewTest:     interviewTest,
		benchCall:         `solutions.ParseHostPort("localhost:8080")`,
	}
}

func databaseSpec(m module) topicSpec {
	body := `package solutions

import "strings"

// BuildSelect builds a parameterized SELECT query.
func BuildSelect(table string, columns []string, where string) string {
	q := "SELECT " + strings.Join(columns, ", ") + " FROM " + table
	if where != "" {
		q += " WHERE " + where
	}
	return q
}

// QuoteIdentifier escapes a SQL identifier (simplified).
func QuoteIdentifier(name string) string {
	return "\"" + strings.ReplaceAll(name, "\"", "\"\"") + "\""
}`
	test := testPackage(m, `BuildSelect("users",[]string{"id","name"},"id=$1")`, `"SELECT id, name FROM users WHERE id=$1"`, `QuoteIdentifier("order")`, `"\"order\""`)
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — SQL, pooling, transactions.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Design repository pattern with transactions",
		interviewSolution: repositoryBody(),
		interviewTest:     testRepository(m),
		benchCall:         `BuildSelect("t",[]string{"a"},"")`,
	}
}

func securitySpec(m module) topicSpec {
	if m.slug == "96-jwt" {
		return jwtExerciseSpec(m)
	}
	body := `package solutions

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPasswordSHA256 hashes password (use bcrypt/argon2 in production).
func HashPasswordSHA256(password string) string {
	h := sha256.Sum256([]byte(password))
	return hex.EncodeToString(h[:])
}

// ConstantTimeCompare compares strings in constant time (simplified).
func ConstantTimeCompare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}`
	test := testPackage(m, `len(HashPasswordSHA256("secret"))`, "64", `ConstantTimeCompare("abc","abc")`, "true")
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — auth, crypto, secure coding.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Implement JWT claims validation",
		interviewSolution: jwtClaimsBody(),
		interviewTest:     testJWTClaims(m),
		benchCall:         `HashPasswordSHA256("bench")`,
	}
}

func testingSpec(m module) topicSpec {
	body := `package solutions

// Add returns sum of two integers (trivial function to test).
func Add(a, b int) int { return a + b }

// IsEven reports whether n is even.
func IsEven(n int) bool { return n%2 == 0 }`
	test := testPackage(m, `Add(2,3)`, "5", `IsEven(4)`, "true")
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — unit, integration, fuzz, benchmarks.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Write table-driven tests with subtests",
		interviewSolution: body,
		interviewTest:     test,
		benchCall:         `Add(1,2)`,
	}
}

func architectureSpec(m module) topicSpec {
	var body, test, problem, interview, iTest string
	switch m.slug {
	case "111-clean-architecture":
		problem = "Separate handler, use case, and repository layers"
		body = cleanArchBody()
		test = testCleanArch(m)
		interview = cleanArchBody()
		iTest = testCleanArch(m)
	case "112-ddd":
		problem = "Model Order aggregate with domain invariants"
		body = dddOrderBody()
		test = testDDDOrder(m)
		interview = dddInterviewBody()
		iTest = testDDDInterview(m)
	case "113-cqrs":
		problem = "Separate command and query handlers"
		body = cqrsBody()
		test = testCQRS(m)
		interview = cqrsBody()
		iTest = testCQRS(m)
	case "114-event-driven":
		problem = "Publish domain events from aggregate"
		body = eventBody()
		test = testEvent(m)
		interview = eventBody()
		iTest = testEvent(m)
	case "118-repository-pattern":
		problem = "Generic repository interface with in-memory impl"
		body = repositoryBody()
		test = testRepository(m)
		interview = repositoryBody()
		iTest = testRepository(m)
	default:
		problem = fmt.Sprintf("Apply %s layering to a user service", m.title)
		body = hexagonalBody()
		test = testHexagonal(m)
		interview = hexagonalBody()
		iTest = testHexagonal(m)
	}
	bench := `GetUserEmail(nil, "1")`
	switch m.slug {
	case "112-ddd":
		bench = `_, _ = NewOrder("o1", []string{"item"})`
	case "113-cqrs":
		bench = `s := NewUserStore(); s.Handle(CreateUserCommand{ID: "1", Email: "a@b.com"})`
	case "114-event-driven":
		bench = `o := &OrderWithEvents{ID: "o1"}; o.Ship()`
	}
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — production architecture patterns.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  problem,
		interviewSolution: interview,
		interviewTest:     iTest,
		benchCall:         bench,
	}
}

func observabilitySpec(m module) topicSpec {
	body := `package solutions

import "fmt"

// MetricPoint represents a single metric sample.
type MetricPoint struct {
	Name  string
	Value float64
	Tags  map[string]string
}

// FormatPrometheus formats a counter metric line.
func FormatPrometheus(name string, value float64) string {
	return fmt.Sprintf("%s %.2f", name, value)
}`
	test := testPackage(m, `FormatPrometheus("http_requests_total", 42)`, `"http_requests_total 42.00"`, ``, ``)
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — logs, metrics, traces.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Design RED/USE metrics for a service",
		interviewSolution: body,
		interviewTest:     test,
		benchCall:         `FormatPrometheus("bench", 1)`,
	}
}

func cloudSpec(m module) topicSpec {
	body := `package solutions

// HealthStatus represents probe result.
type HealthStatus struct {
	OK     bool
	Reason string
}

// Liveness returns alive status.
func Liveness() HealthStatus { return HealthStatus{OK: true} }

// Readiness checks dependencies (stub).
func Readiness(dbOK, cacheOK bool) HealthStatus {
	if dbOK && cacheOK {
		return HealthStatus{OK: true}
	}
	return HealthStatus{OK: false, Reason: "dependency unavailable"}
}`
	test := testPackage(m, `Liveness().OK`, "true", `Readiness(true,true).OK`, "true")
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — containers, K8s, CI/CD, IaC.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Design Kubernetes liveness/readiness probes",
		interviewSolution: body,
		interviewTest:     test,
		benchCall:         `Liveness()`,
	}
}

func messagingSpec(m module) topicSpec {
	body := `package solutions

import "time"

// Message represents an event on a bus.
type Message struct {
	Topic     string
	Key       string
	Payload   []byte
	Timestamp time.Time
}

// NewMessage creates a message with current timestamp.
func NewMessage(topic, key string, payload []byte) Message {
	return Message{Topic: topic, Key: key, Payload: payload, Timestamp: time.Now()}
}`
	test := testPackage(m, `NewMessage("orders","k",[]byte("x")).Topic`, `"orders"`, ``, ``)
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — pub/sub, streaming, queues.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Design idempotent Kafka consumer",
		interviewSolution: idempotentConsumerBody(),
		interviewTest:     testIdempotentConsumer(m),
		benchCall:         `NewMessage("t","k",nil)`,
	}
}

func systemDesignSpec(m module) topicSpec {
	var body, test, problem, interview, interviewTest string
	switch m.slug {
	case "144-caching":
		body = lruCacheBody()
		test = testLRU(m)
		problem = "Implement LRU cache O(1) get/put"
		interview = lruInterviewBody()
		interviewTest = testLRUInterview(m)
	case "145-rate-limiting":
		body = tokenBucketBody()
		test = testTokenBucket(m)
		problem = "Token bucket rate limiter"
		interview = rateLimitSlidingWindowBody()
		interviewTest = testRateLimitSlidingWindow(m)
	case "146-circuit-breaker":
		body = circuitBreakerBody()
		test = testCircuitBreaker(m)
		problem = "Circuit breaker with open/half-open/closed"
		interview = circuitBreakerInterviewBody()
		interviewTest = testCircuitBreakerInterview(m)
	default:
		body = urlShortenerBody()
		test = testURLShortener(m)
		problem = "Design URL shortener encode/decode"
		interview = urlShortenerInterviewBody()
		interviewTest = testURLShortenerInterview(m)
	}
	bench := `solutions.NewTokenBucket(10).Allow()`
	if m.slug == "144-caching" {
		bench = `c := solutions.NewLRU(2); c.Put("a", 1)`
	} else if m.slug == "146-circuit-breaker" {
		bench = `cb := solutions.NewCircuitBreaker(3); cb.Call(func() error { return nil })`
	} else if m.slug != "145-rate-limiting" {
		bench = `solutions.Encode(3)`
	}
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — scalable system components.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  problem,
		interviewSolution: interview,
		interviewTest:     interviewTest,
		benchCall:         bench,
	}
}

func metaSpec(m module) topicSpec {
	body := `package solutions

// ModuleProgress tracks learning progress.
type ModuleProgress struct {
	Module  string
	Percent int
}

// Completed reports whether module is done.
func (p ModuleProgress) Completed() bool { return p.Percent >= 100 }`
	test := testPackage(m, `ModuleProgress{"x",100}.Completed()`, "true", `ModuleProgress{"x",50}.Completed()`, "false")
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — roadmap and reference material.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Plan a learning path for Go backend engineering",
		interviewSolution: body,
		interviewTest:     test,
		benchCall:         `ModuleProgress{"b",100}.Completed()`,
	}
}

func ioSpec(m module) topicSpec {
	body := `package solutions

import "strings"

// Lines splits text into non-empty trimmed lines.
func Lines(text string) []string {
	raw := strings.Split(text, "\n")
	out := make([]string, 0, len(raw))
	for _, l := range raw {
		l = strings.TrimSpace(l)
		if l != "" {
			out = append(out, l)
		}
	}
	return out
}`
	test := testPackage(m, `len(Lines("a\n\nb"))`, "2", `len(Lines(""))`, "0")
	return topicSpec{
		summary:           fmt.Sprintf("**%s** — file I/O and serialization.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Stream-parse large JSON without loading entire file",
		interviewSolution: body,
		interviewTest:     test,
		benchCall:         `Lines("a\nb\nc")`,
	}
}

var overrides = map[string]topicSpec{}
