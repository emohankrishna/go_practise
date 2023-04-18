package main

import "fmt"

// Detect a cycle in undirected graph
type Tuple struct {
	dst, src int
}

func detectCycleWithBFS(g map[int][]int, visited []bool, src int) bool {
	var q []Tuple = []Tuple{{dst: src, src: -1}}
	visited[src] = true
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		adjList := g[node.dst]
		for _, v := range adjList {
			if !visited[v] {
				visited[v] = true
				q = append(q, Tuple{
					v, node.dst,
				})

			} else if v != node.src {
				return true
			}
		}
	}
	return false
}

func detectCycleWithDFS(g map[int][]int, visited []bool, src int) bool {
	var q []Tuple = []Tuple{{dst: src, src: -1}}
	visited[src] = true
	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		adjList := g[node.dst]
		for _, v := range adjList {
			if !visited[v] {
				visited[v] = true
				q = append(q, Tuple{
					v, node.dst,
				})

			} else if v != node.src {
				return true
			}
		}
	}
	return false
}

func isContainCycle(g map[int][]int, v int) bool {
	visited := make([]bool, len(g)+1)
	for i := range g {
		if !visited[i] {
			return detectCycleWithDFS(g, visited, i)
		}
	}
	return false
}
func main() {
	g := map[int][]int{
		1: {2, 4},
		2: {3, 1},
		3: {2, 6},
		4: {7, 5, 1},
		5: {6, 4},
		6: {3, 5},
		7: {4},
	}
	fmt.Println(isContainCycle(g, 1))
}
