package solutions

// InterviewChallengeSolution returns minimum of two ints.
func InterviewChallengeSolution(a, b int) int {
	if a < b {
		return a
	}
	return b
}