package solutions

type UnionFind struct{ parent []int }

func NewUF(n int) *UnionFind {
	p := make([]int, n)
	for i := range p { p[i] = i }
	return &UnionFind{p}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x { uf.parent[x] = uf.Find(uf.parent[x]) }
	return uf.parent[x]
}

func (uf *UnionFind) Union(a, b int) { pa, pb := uf.Find(a), uf.Find(b); uf.parent[pa] = pb }