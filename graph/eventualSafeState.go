package main

import "fmt"

func topologicalSort(V int, adj map[int][]int) []int {
	inDegree := make([]int, V)
	for _, val := range adj {
		for _, v := range val {
			inDegree[v] += 1
		}
	}
	q := []int{}
	for k, v := range inDegree {
		if v == 0 {
			q = append(q, k)
		}
	}
	topo := []int{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		// Add the node to result
		topo = append(topo, node)
		for _, v := range adj[node] {
			inDegree[v] -= 1
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return topo
}

func eventualSafeNodes(graph [][]int) []int {
	m := len(graph)
	graphCom := make(map[int][]int)
	// Create new graph where we are reversing the edges
	for i := 0; i < m; i++ {
		graphCom[i] = []int{}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < len(graph[i]); j++ {
			graphCom[graph[i][j]] = append(graphCom[graph[i][j]], i)
		}
	}
	// apply the topological sort eventually you will get the safe states

	return topologicalSort(m, graphCom)
}

func main() {
	g := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	fmt.Println(eventualSafeNodes(g))
}
