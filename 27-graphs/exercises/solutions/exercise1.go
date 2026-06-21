package solutions

// Graph adjacency list.
type Graph map[int][]int

func (g Graph) BFS(start int) []int {
	visited := map[int]bool{start: true}
	q := []int{start}
	order := []int{start}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, nei := range g[v] {
			if !visited[nei] {
				visited[nei] = true
				q = append(q, nei)
				order = append(order, nei)
			}
		}
	}
	return order
}