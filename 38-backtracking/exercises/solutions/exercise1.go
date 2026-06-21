package solutions

func Permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(nums) {
			cp := append([]int(nil), nums...)
			res = append(res, cp)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtrack(0)
	return res
}