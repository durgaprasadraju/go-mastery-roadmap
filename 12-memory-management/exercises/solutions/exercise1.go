package solutions

// SharesBacking reports whether a and b share the same backing array.
func SharesBacking(a, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	return &a[0] == &b[0]
}