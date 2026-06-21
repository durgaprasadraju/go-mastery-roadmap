package solutions

import "sync"

// ParallelSum sums slice using goroutines.
func ParallelSum(nums []int, workers int) int {
	if workers < 1 { workers = 1 }
	chunk := (len(nums) + workers - 1) / workers
	var wg sync.WaitGroup
	partial := make([]int, workers)
	for w := 0; w < workers; w++ {
		start := w * chunk
		if start >= len(nums) { break }
		end := start + chunk
		if end > len(nums) { end = len(nums) }
		wg.Add(1)
		go func(idx, lo, hi int) {
			defer wg.Done()
			for i := lo; i < hi; i++ { partial[idx] += nums[i] }
		}(w, start, end)
	}
	wg.Wait()
	total := 0
	for _, p := range partial { total += p }
	return total
}