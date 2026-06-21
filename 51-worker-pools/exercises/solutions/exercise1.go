package solutions

// RunPool processes n jobs with w workers.
func RunPool(jobs, workers int) int {
	ch := make(chan int, jobs)
	for i := 0; i < jobs; i++ { ch <- i }
	close(ch)
	done := make(chan struct{}, workers)
	for w := 0; w < workers; w++ {
		go func() {
			for range ch {}
			done <- struct{}{}
		}()
	}
	for w := 0; w < workers; w++ { <-done }
	return jobs
}