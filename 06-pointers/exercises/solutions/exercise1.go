package solutions

// Swap exchanges values at a and b.
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// IntPtr returns a pointer to v.
func IntPtr(v int) *int { return &v }