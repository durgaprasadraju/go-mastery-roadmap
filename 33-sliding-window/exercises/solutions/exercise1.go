package solutions

func MaxSumSubarrayK(arr []int, k int) int {
	if len(arr) < k { return 0 }
	sum := 0
	for i := 0; i < k; i++ { sum += arr[i] }
	max := sum
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if sum > max { max = sum }
	}
	return max
}