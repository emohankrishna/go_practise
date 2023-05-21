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
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, v := range prerequisites {
		list := graph[v[1]]
		list = append(list, v[0])
		graph[v[1]] = list
	}
	return len(topologicalSort(numCourses, graph)) == numCourses

}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
