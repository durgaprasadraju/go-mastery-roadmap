package solutions

import "strconv"

// FizzBuzz returns FizzBuzz strings for 1..n.
func FizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		out[i-1] = s
	}
	return out
}