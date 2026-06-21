package solutions

// RuneLen returns number of runes in string.
func RuneLen(s string) int {
	return len([]rune(s))
}