package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to   int
	cost int
}

func bfs(graph [][]Edge, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	queue := make([]int, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, e := range graph[u] {
			if dist[e.to] > dist[u]+e.cost {
				dist[e.to] = dist[u] + e.cost
				queue = append(queue, e.to)
			}
		}
	}

	return dist
}

func main() {
	graph := [][]Edge{
		{{to: 1, cost: 4}, {to: 2, cost: 1}},
		{{to: 0, cost: 4}, {to: 2, cost: 2}, {to: 3, cost: 5}},
		{{to: 0, cost: 1}, {to: 1, cost: 2}, {to: 3, cost: 1}},
		{{to: 1, cost: 5}, {to: 2, cost: 1}},
	}
	start := 0

	dist := bfs(graph, start)

	fmt.Println("Shortest distances from vertex", start, "to all vertices:")
	for i, d := range dist {
		fmt.Printf("%d: %d\n", i, d)
	}
}
