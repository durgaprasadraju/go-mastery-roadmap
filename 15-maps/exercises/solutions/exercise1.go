package solutions

// WordCount counts word frequencies.
func WordCount(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	return m
}