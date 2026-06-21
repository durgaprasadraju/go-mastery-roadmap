package solutions

import "strings"

// ParseHostPort splits host:port (simplified).
func ParseHostPort(addr string) (host, port string, ok bool) {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

// BuildURL builds http URL from parts.
func BuildURL(host, path string) string {
	return "http://" + host + path
}