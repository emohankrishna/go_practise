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
	fmt.Println(inDegree, q)
	fmt.Println()
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
	fmt.Println(inDegree, q)
	return topo
}
func main() {
	g := map[int][]int{
		0: {},
		1: {},
		2: {3},
		3: {1},
		4: {0, 1},
		5: {0, 2},
	}
	fmt.Println(topologicalSort(len(g), g)) // [4 5 0 2 3 1]
}
