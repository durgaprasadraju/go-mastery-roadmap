// Package graphs implements adjacency-list graph with BFS and DFS.
package graphs

import "errors"

var ErrNotFound = errors.New("vertex not found")

// Graph is a directed/undirected adjacency-list graph.
type Graph struct {
	adj     map[int][]int
	undir   bool
	vertices int
}

// New creates a graph. Set undirected=true for bidirectional edges.
func New(undirected bool) *Graph {
	return &Graph{adj: make(map[int][]int), undir: undirected}
}

// AddVertex adds a vertex.
func (g *Graph) AddVertex(v int) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = make([]int, 0)
		g.vertices++
	}
}

// AddEdge adds an edge from u to v.
func (g *Graph) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)
	g.adj[u] = append(g.adj[u], v)
	if g.undir {
		g.adj[v] = append(g.adj[v], u)
	}
}

// BFS returns vertices in breadth-first order from start. O(V+E).
func (g *Graph) BFS(start int) ([]int, error) {
	if _, ok := g.adj[start]; !ok {
		return nil, ErrNotFound
	}
	visited := make(map[int]bool)
	queue := []int{start}
	order := make([]int, 0)
	visited[start] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, nei := range g.adj[v] {
			if !visited[nei] {
				visited[nei] = true
				queue = append(queue, nei)
			}
		}
	}
	return order, nil
}

// DFS returns vertices in depth-first order. O(V+E).
func (g *Graph) DFS(start int) ([]int, error) {
	if _, ok := g.adj[start]; !ok {
		return nil, ErrNotFound
	}
	visited := make(map[int]bool)
	order := make([]int, 0)
	var walk func(int)
	walk = func(v int) {
		visited[v] = true
		order = append(order, v)
		for _, nei := range g.adj[v] {
			if !visited[nei] {
				walk(nei)
			}
		}
	}
	walk(start)
	return order, nil
}

// HasPath reports whether a path exists from src to dst.
func (g *Graph) HasPath(src, dst int) bool {
	order, err := g.BFS(src)
	if err != nil {
		return false
	}
	for _, v := range order {
		if v == dst {
			return true
		}
	}
	return false
}

// VertexCount returns the number of vertices.
func (g *Graph) VertexCount() int { return g.vertices }
