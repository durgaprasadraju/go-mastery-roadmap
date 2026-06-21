package solutions

func TwoSumSorted(arr []int, target int) [2]int {
	i, j := 0, len(arr)-1
	for i < j {
		s := arr[i] + arr[j]
		if s == target { return [2]int{i,j} }
		if s < target { i++ } else { j-- }
	}
	return [2]int{-1,-1}
}