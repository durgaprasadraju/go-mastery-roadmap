package solutions

// TwoSum returns indices of two numbers summing to target.
func TwoSum(nums []int, target int) [2]int {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return [2]int{j, i}
		}
		seen[n] = i
	}
	return [2]int{-1, -1}
}