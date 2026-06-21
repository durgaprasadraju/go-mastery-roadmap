package solutions

// LIS returns length of longest increasing subsequence.
func LIS(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	tails := make([]int, 0, len(arr))
	for _, v := range arr {
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if tails[mid] < v {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, v)
		} else {
			tails[lo] = v
		}
	}
	return len(tails)
}