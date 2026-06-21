package solutions

// InterviewChallengeSolution returns the maximum sum of any contiguous subarray (Kadane's).
// Demonstrates Tries interview pattern. O(n) time, O(1) space.
func InterviewChallengeSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxCurrent, maxGlobal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxCurrent+nums[i] > nums[i] {
			maxCurrent += nums[i]
		} else {
			maxCurrent = nums[i]
		}
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}
	return maxGlobal
}
