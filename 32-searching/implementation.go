// Package searching implements linear, binary, and ternary search algorithms.
package searching

// LinearSearch finds target in arr. O(n).
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// BinarySearch finds target in sorted arr. O(log n).
func BinarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// TernarySearch finds target in unimodal sorted arr. O(log3 n).
func TernarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid1 := lo + (hi-lo)/3
		mid2 := hi - (hi-lo)/3
		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}
		if target < arr[mid1] {
			hi = mid1 - 1
		} else if target > arr[mid2] {
			lo = mid2 + 1
		} else {
			lo = mid1 + 1
			hi = mid2 - 1
		}
	}
	return -1
}
