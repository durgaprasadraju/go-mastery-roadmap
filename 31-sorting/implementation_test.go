package sorting_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/31-sorting"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3}
	sorting.MergeSort(arr)
	if !sorting.IsSorted(arr) {
		t.Fatalf("not sorted: %v", arr)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3}
	sorting.QuickSort(arr)
	if !sorting.IsSorted(arr) {
		t.Fatalf("not sorted: %v", arr)
	}
}

func TestAllSorts(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3}
	sorters := []struct {
		name string
		fn   func([]int)
	}{
		{"bubble", sorting.BubbleSort},
		{"selection", sorting.SelectionSort},
		{"insertion", sorting.InsertionSort},
		{"merge", sorting.MergeSort},
		{"quick", sorting.QuickSort},
	}
	for _, s := range sorters {
		arr := append([]int(nil), original...)
		s.fn(arr)
		if !sorting.IsSorted(arr) {
			t.Errorf("%s: not sorted %v", s.name, arr)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 1000 - i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := append([]int(nil), data...)
		sorting.QuickSort(arr)
	}
}
