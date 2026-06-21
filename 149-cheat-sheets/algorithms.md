# Algorithms Cheat Sheet

## Sorting

| Algorithm | Best | Average | Worst | Space | Stable |
|-----------|------|---------|-------|-------|--------|
| Bubble | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Selection | O(n²) | O(n²) | O(n²) | O(1) | No |
| Insertion | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Merge | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| Quick | O(n log n) | O(n log n) | O(n²) | O(log n) | No |
| Heap | O(n log n) | O(n log n) | O(n log n) | O(1) | No |
| Counting | O(n+k) | O(n+k) | O(n+k) | O(k) | Yes |
| Radix | O(nk) | O(nk) | O(nk) | O(n+k) | Yes |

## Searching

| Algorithm | Time | Requirement |
|-----------|------|-------------|
| Linear | O(n) | Any array |
| Binary | O(log n) | Sorted array |
| Ternary | O(log3 n) | Unimodal sorted |

## Graph Algorithms

| Algorithm | Time | Use Case |
|-----------|------|----------|
| BFS | O(V+E) | Shortest path (unweighted) |
| DFS | O(V+E) | Cycle detection, topological sort |
| Dijkstra | O((V+E) log V) | Shortest path (weighted) |
| Bellman-Ford | O(VE) | Negative weights |
| Floyd-Warshall | O(V³) | All-pairs shortest path |
| Prim/Kruskal | O(E log V) | Minimum spanning tree |
| Topological Sort | O(V+E) | Dependency ordering |

## Dynamic Programming

| Problem | Time | Space |
|---------|------|-------|
| Knapsack 0/1 | O(nW) | O(nW) |
| LIS | O(n²) / O(n log n) | O(n) |
| LCS | O(mn) | O(mn) |
| Coin Change | O(n×amount) | O(amount) |
| Matrix Chain | O(n³) | O(n²) |

## String Algorithms

| Algorithm | Time | Use Case |
|-----------|------|----------|
| KMP | O(n+m) | Pattern matching |
| Rabin-Karp | O(n+m) avg | Multiple pattern search |
| Boyer-Moore | O(n/m) best | Large alphabet |
| Trie Search | O(m) | Prefix matching |

## Go stdlib

```go
sort.Ints(arr)           // O(n log n)
sort.Search(n, func(i int) bool { return arr[i] >= target })
```
