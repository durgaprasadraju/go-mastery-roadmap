package solutions

import "strings"

// InterviewChallengeSolution validates JWT has three base64url segments.
func InterviewChallengeSolution(token string) bool {
	parts := strings.Split(token, ".")
	return len(parts) == 3 && parts[0] != "" && parts[1] != ""
}