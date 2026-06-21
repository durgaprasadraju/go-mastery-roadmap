package solutions

// FindMax returns maximum in array.
func FindMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, v := range arr[1:] {
		if v > m {
			m = v
		}
	}
	return m
}