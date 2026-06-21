// Package segmenttree implements a segment tree for range sum queries.
package segmenttree

// SegmentTree supports range sum queries and point updates in O(log n).
type SegmentTree struct {
	tree []int
	n    int
}

// New builds a segment tree from data.
func New(data []int) *SegmentTree {
	n := len(data)
	st := &SegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	if n > 0 {
		st.build(data, 0, 0, n-1)
	}
	return st
}

func (st *SegmentTree) build(data []int, node, start, end int) {
	if start == end {
		st.tree[node] = data[start]
		return
	}
	mid := (start + end) / 2
	left := 2*node + 1
	right := 2*node + 2
	st.build(data, left, start, mid)
	st.build(data, right, mid+1, end)
	st.tree[node] = st.tree[left] + st.tree[right]
}

// Query returns sum in range [l, r]. O(log n).
func (st *SegmentTree) Query(l, r int) int {
	if st.n == 0 {
		return 0
	}
	return st.query(0, 0, st.n-1, l, r)
}

func (st *SegmentTree) query(node, start, end, l, r int) int {
	if r < start || end < l {
		return 0
	}
	if l <= start && end <= r {
		return st.tree[node]
	}
	mid := (start + end) / 2
	left := st.query(2*node+1, start, mid, l, r)
	right := st.query(2*node+2, mid+1, end, l, r)
	return left + right
}

// Update sets data[idx] = val. O(log n).
func (st *SegmentTree) Update(idx, val int) {
	if st.n == 0 {
		return
	}
	st.update(0, 0, st.n-1, idx, val)
}

func (st *SegmentTree) update(node, start, end, idx, val int) {
	if start == end {
		st.tree[node] = val
		return
	}
	mid := (start + end) / 2
	if idx <= mid {
		st.update(2*node+1, start, mid, idx, val)
	} else {
		st.update(2*node+2, mid+1, end, idx, val)
	}
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}
