package solutions

func MaxActivities(ends []int) int {
	count, last := 0, -1
	for i, end := range ends {
		if i == 0 || end >= last {
			count++
			last = end
		}
	}
	return count
}