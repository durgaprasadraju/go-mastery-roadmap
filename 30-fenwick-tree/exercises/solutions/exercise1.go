package solutions

// RangeSum returns sum of slice.
func RangeSum(arr []int) int {
	s := 0
	for _, v := range arr { s += v }
	return s
}