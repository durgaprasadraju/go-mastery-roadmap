package solutions

import "strings"

// InterviewChallengeSolution matches HTTP method and path pattern.
func InterviewChallengeSolution(method, path, wantMethod, wantPrefix string) bool {
	return strings.EqualFold(method, wantMethod) && strings.HasPrefix(path, wantPrefix)
}