package solutions

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	left := append([]int(nil), arr[:mid]...)
	right := append([]int(nil), arr[mid:]...)
	MergeSort(left)
	MergeSort(right)
	merge(arr, left, right)
}

func merge(dst, left, right []int) {
	i,j,k := 0,0,0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] { dst[k]=left[i]; i++ } else { dst[k]=right[j]; j++ }
		k++
	}
	for i < len(left) { dst[k]=left[i]; i++; k++ }
	for j < len(right) { dst[k]=right[j]; j++; k++ }
}