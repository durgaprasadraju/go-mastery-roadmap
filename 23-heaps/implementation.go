// Package heaps implements a min-heap priority queue from scratch.
package heaps

import "errors"

var ErrEmpty = errors.New("heap is empty")

// MinHeap is a binary min-heap stored in a slice.
type MinHeap struct {
	data []int
}

// NewMinHeap creates an empty min-heap.
func NewMinHeap() *MinHeap {
	return &MinHeap{data: make([]int, 0)}
}

// Size returns element count.
func (h *MinHeap) Size() int { return len(h.data) }

// IsEmpty reports whether the heap is empty.
func (h *MinHeap) IsEmpty() bool { return len(h.data) == 0 }

// Push inserts a value. O(log n).
func (h *MinHeap) Push(v int) {
	h.data = append(h.data, v)
	h.siftUp(len(h.data) - 1)
}

// Pop removes and returns the minimum. O(log n).
func (h *MinHeap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, ErrEmpty
	}
	min := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if len(h.data) > 0 {
		h.siftDown(0)
	}
	return min, nil
}

// Peek returns the minimum without removing. O(1).
func (h *MinHeap) Peek() (int, error) {
	if len(h.data) == 0 {
		return 0, ErrEmpty
	}
	return h.data[0], nil
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] >= h.data[parent] {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *MinHeap) siftDown(i int) {
	n := len(h.data)
	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2
		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

// HeapSort sorts in-place using heap sort. O(n log n).
func HeapSort(arr []int) {
	h := &MinHeap{data: make([]int, len(arr))}
	copy(h.data, arr)
	// Build max-heap variant for in-place sort — simplified: collect via pop
	result := make([]int, len(arr))
	for i := range arr {
		h.Push(arr[i])
	}
	for i := range result {
		v, _ := h.Pop()
		result[i] = v
	}
	copy(arr, result)
}
