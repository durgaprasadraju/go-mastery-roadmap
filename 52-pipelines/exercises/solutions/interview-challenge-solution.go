package solutions

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
}