// Package fenwick implements a Fenwick tree (Binary Indexed Tree) for prefix sums.
package fenwick

// FenwickTree supports prefix sum and point update in O(log n).
type FenwickTree struct {
	tree []int
	n    int
}

// New creates a Fenwick tree from data.
func New(data []int) *FenwickTree {
	n := len(data)
	ft := &FenwickTree{tree: make([]int, n+1), n: n}
	for i, v := range data {
		ft.Add(i, v)
	}
	return ft
}

// Add adds delta to index i (0-based).
func (ft *FenwickTree) Add(i, delta int) {
	for idx := i + 1; idx <= ft.n; idx += idx & -idx {
		ft.tree[idx] += delta
	}
}

// PrefixSum returns sum of [0..i]. O(log n).
func (ft *FenwickTree) PrefixSum(i int) int {
	sum := 0
	for idx := i + 1; idx > 0; idx -= idx & -idx {
		sum += ft.tree[idx]
	}
	return sum
}

// RangeSum returns sum of [l..r]. O(log n).
func (ft *FenwickTree) RangeSum(l, r int) int {
	if l == 0 {
		return ft.PrefixSum(r)
	}
	return ft.PrefixSum(r) - ft.PrefixSum(l-1)
}
