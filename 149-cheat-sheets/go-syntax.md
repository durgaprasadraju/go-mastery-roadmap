# Go Syntax Cheat Sheet

## Types

| Type | Zero Value | Example |
|------|-----------|---------|
| `bool` | `false` | `var ok bool` |
| `int` | `0` | `x := 42` |
| `string` | `""` | `s := "hello"` |
| `*T` | `nil` | `var p *int` |
| `[]T` | `nil` | `sl := []int{1,2,3}` |
| `map[K]V` | `nil` | `m := make(map[string]int)` |
| `chan T` | `nil` | `ch := make(chan int, 10)` |

## Control Flow

```go
if x > 0 { } else if x < 0 { } else { }
for i := 0; i < n; i++ { }
for k, v := range m { }
switch v := x.(type) { case int: _ = v }
```

## Error Handling

```go
if err != nil {
    return fmt.Errorf("context: %w", err)
}
```

## Interfaces & Type Assertions

```go
var r io.Reader = os.Stdin
if w, ok := r.(io.Writer); ok { w.Write([]byte("hi")) }
```

## Struct Tags

```go
type User struct {
    ID    int    `json:"id" db:"id"`
    Email string `json:"email" validate:"required,email"`
}
```

## Generics (Go 1.18+)

```go
func Map[T, U any](s []T, fn func(T) U) []U {
    out := make([]U, len(s))
    for i, v := range s { out[i] = fn(v) }
    return out
}
```

## Testing

```go
func TestAdd(t *testing.T) {
    tests := []struct{ a, b, want int }{
        {1, 2, 3}, {0, 0, 0},
    }
    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("got %d want %d", got, tt.want)
        }
    }
}
```
