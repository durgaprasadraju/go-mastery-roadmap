package solutions

// Fib returns nth Fibonacci number with memoization.
func Fib(n int) int {
	memo := map[int]int{0: 0, 1: 1}
	var f func(int) int
	f = func(k int) int {
		if v, ok := memo[k]; ok {
			return v
		}
		memo[k] = f(k-1) + f(k-2)
		return memo[k]
	}
	return f(n)
}