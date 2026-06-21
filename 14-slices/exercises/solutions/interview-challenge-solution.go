package solutions

// ReverseInPlace reverses slice in O(n) time, O(1) space.
func ReverseInPlace(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}