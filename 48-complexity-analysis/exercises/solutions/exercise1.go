package solutions

// CountOperations returns op count for linear scan of n elements.
func CountOperations(n int) int { return n }

// CountNested returns op count for naive O(n^2) nested loop.
func CountNested(n int) int { return n * n }

// ClassifyGrowth returns big-O label for common growth rates.
func ClassifyGrowth(n int, kind string) string {
	switch kind {
	case "constant":
		return "O(1)"
	case "linear":
		return "O(n)"
	case "quadratic":
		return "O(n^2)"
	case "log":
		return "O(log n)"
	default:
		return "unknown"
	}
}