# Concurrency Cheat Sheet

## Goroutines

```go
go func() { fmt.Println("async") }()
```

## Channels

```go
ch := make(chan int)       // unbuffered
ch := make(chan int, 100)  // buffered
ch <- 42                   // send
v := <-ch                  // receive
close(ch)                  // close (sender only)
```

## Select

```go
select {
case v := <-ch1:
    _ = v
case ch2 <- 42:
case <-ctx.Done():
    return ctx.Err()
default:
    // non-blocking
}
```

## sync Package

```go
var mu sync.Mutex
mu.Lock(); defer mu.Unlock()

var rw sync.RWMutex
rw.RLock(); defer rw.RUnlock()

var wg sync.WaitGroup
wg.Add(1); go func() { defer wg.Done() }(); wg.Wait()

var once sync.Once
once.Do(func() { /* init */ })
```

## Atomic

```go
var count atomic.Int64
count.Add(1)
count.Load()
```

## Context

```go
ctx, cancel := context.WithCancel(parent)
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
ctx = context.WithValue(ctx, key, value)
```

## Patterns

| Pattern | Use Case |
|---------|----------|
| Worker Pool | Bounded parallelism |
| Pipeline | Sequential processing stages |
| Fan-Out/Fan-In | Parallel map-reduce |
| Semaphore | Rate limiting |
| errgroup | Parallel tasks with error propagation |

## Memory Model Rules

1. A send on a channel happens before the corresponding receive
2. Unlock happens before any later Lock on the same mutex
3. Once.Do happens before any call to the function returns

## Debugging

```bash
go test -race ./...
go run -race main.go
```
