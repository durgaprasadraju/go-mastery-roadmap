// Package unionfind implements disjoint-set (Union-Find) with path compression and union by rank.
package unionfind

// UnionFind tracks connected components.
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

// New creates a Union-Find structure for n elements (0..n-1).
func New(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank, count: n}
}

// Find returns the representative of set containing x. O(α(n)) amortized.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

// Union merges sets containing x and y. O(α(n)) amortized.
func (uf *UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return false
	}
	// union by rank
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	uf.count--
	return true
}

// Connected reports whether x and y are in the same set.
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count returns the number of disjoint sets.
func (uf *UnionFind) Count() int { return uf.count }
