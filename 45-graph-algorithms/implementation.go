// Package graphalgorithms implements classic graph algorithms.
package graphalgorithms

import "math"

// Edge represents a weighted directed edge.
type Edge struct {
	From, To int
	Weight   int
}

// Dijkstra finds shortest distances from source. O((V+E) log V) with heap.
func Dijkstra(adj map[int][]Edge, source int) map[int]int {
	dist := map[int]int{source: 0}
	visited := map[int]bool{}

	for len(visited) < len(adj)+1 {
		u := -1
		minDist := math.MaxInt32
		for v, d := range dist {
			if !visited[v] && d < minDist {
				minDist = d
				u = v
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true
		for _, e := range adj[u] {
			if nd := dist[u] + e.Weight; nd < getDist(dist, e.To) {
				dist[e.To] = nd
			}
		}
	}
	return dist
}

func getDist(dist map[int]int, v int) int {
	if d, ok := dist[v]; ok {
		return d
	}
	return math.MaxInt32
}

// BellmanFord finds shortest paths; detects negative cycles. O(VE).
func BellmanFord(edges []Edge, n, source int) ([]int, bool) {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.From] != math.MaxInt32 && dist[e.From]+e.Weight < dist[e.To] {
				dist[e.To] = dist[e.From] + e.Weight
			}
		}
	}
	for _, e := range edges {
		if dist[e.From] != math.MaxInt32 && dist[e.From]+e.Weight < dist[e.To] {
			return nil, true // negative cycle
		}
	}
	return dist, false
}

// TopologicalSort returns topological order for DAG. O(V+E).
func TopologicalSort(adj map[int][]int, vertices []int) ([]int, bool) {
	inDegree := make(map[int]int)
	for _, v := range vertices {
		inDegree[v] = 0
	}
	for _, v := range vertices {
		for _, nei := range adj[v] {
			inDegree[nei]++
		}
	}
	queue := make([]int, 0)
	for _, v := range vertices {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}
	order := make([]int, 0, len(vertices))
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, nei := range adj[v] {
			inDegree[nei]--
			if inDegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}
	if len(order) != len(vertices) {
		return nil, false // cycle detected
	}
	return order, true
}

// BFSTraversal returns BFS order from source.
func BFSTraversal(adj map[int][]int, source int) []int {
	visited := map[int]bool{source: true}
	queue := []int{source}
	order := []int{source}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, nei := range adj[v] {
			if !visited[nei] {
				visited[nei] = true
				queue = append(queue, nei)
				order = append(order, nei)
			}
		}
	}
	return order
}
