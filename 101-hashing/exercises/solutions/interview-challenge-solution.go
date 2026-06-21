package solutions

import "strings"

// ParseBearer extracts token from Authorization header.
func ParseBearer(header string) (string, bool) {
	const p = "Bearer "
	if !strings.HasPrefix(header, p) { return "", false }
	return strings.TrimSpace(header[len(p):]), true
}