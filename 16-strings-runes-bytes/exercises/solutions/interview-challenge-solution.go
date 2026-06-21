package solutions

import "strings"

// IsPalindrome checks if s is palindrome (case-insensitive).
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}