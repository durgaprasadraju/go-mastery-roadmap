package solutions

// InterviewChallengeSolution decodes short code to id.
func InterviewChallengeSolution(code string) int {
	if code == "" {
		return 0
	}
	return int(code[0] - 'a')
}