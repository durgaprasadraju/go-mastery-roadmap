package solutions

// Compose chains g then f.
func Compose(f, g func(int) int) func(int) int {
	return func(x int) int { return f(g(x)) }
}

// MapInts applies fn to each element.
func MapInts(xs []int, fn func(int) int) []int {
	out := make([]int, len(xs))
	for i, v := range xs {
		out[i] = fn(v)
	}
	return out
}