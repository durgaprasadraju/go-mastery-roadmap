package solutions

// InterviewChallengeSolution reports LRU cache size after puts.
func InterviewChallengeSolution(keys []string, cap int) int {
	c := NewLRU(cap)
	for _, k := range keys {
		c.Put(k, 1)
	}
	return len(c.order)
}