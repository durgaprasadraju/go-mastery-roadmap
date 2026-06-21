package solutions

// InterviewChallengeSolution checks whether order meets minimum item rule.
func InterviewChallengeSolution(itemCount, minItems int) bool {
	return itemCount >= minItems && minItems > 0
}