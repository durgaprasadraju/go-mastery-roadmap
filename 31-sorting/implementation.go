// Package sorting implements classic sorting algorithms from scratch.
package sorting

// BubbleSort sorts in O(n²) time. Educational only — not for production.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// SelectionSort sorts in O(n²) time.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// InsertionSort sorts in O(n²) worst case, O(n) best for nearly sorted.
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// MergeSort sorts in O(n log n) time and O(n) space.
func MergeSort(arr []int) {
	if len(arr) <= 1 {
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
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			dst[k] = left[i]
			i++
		} else {
			dst[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		dst[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		dst[k] = right[j]
		j++
		k++
	}
}

// QuickSort sorts in O(n log n) average, O(n²) worst.
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// CountingSort sorts non-negative integers in O(n+k) time.
func CountingSort(arr []int, maxVal int) {
	count := make([]int, maxVal+1)
	for _, v := range arr {
		count[v]++
	}
	idx := 0
	for v, c := range count {
		for c > 0 {
			arr[idx] = v
			idx++
			c--
		}
	}
}

// IsSorted verifies ascending order.
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
