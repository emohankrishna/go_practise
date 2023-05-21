package main

import (
	"fmt"
	"math"
)

type Edge struct {
	v, cost int
}

func dfs(v int, graph map[int][]Edge, visited []bool, stack []int) []int {
	visited[v] = true
	for _, node := range graph[v] {
		if !visited[node.v] {
			stack = dfs(node.v, graph, visited, stack)
		}
	}
	stack = append(stack, v)
	return stack
}

func shortestPath(N int, edges [][]int, src int) []int {
	adj := make(map[int][]Edge)
	distanceArray := make([]int, N)

	for i := 0; i < N; i++ {
		adj[i] = []Edge{}
		distanceArray[i] = math.MaxInt32
	}
	distanceArray[src] = 0
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], Edge{v[1], v[2]})
	}
	visited := make([]bool, N)
	stack := []int{}
	for i := 0; i < N; i++ {
		if !visited[i] {
			stack = dfs(i, adj, visited, stack)
		}
	}
	fmt.Println(stack)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top == src {
			break
		}
		stack = stack[:len(stack)-1]
	}
	fmt.Println(stack)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		l := adj[top]
		for _, k := range l {
			if distanceArray[top]+k.cost < distanceArray[k.v] {
				distanceArray[k.v] = k.cost + distanceArray[top]
			}
		}

	}
	for i := 0; i < len(distanceArray); i++ {
		if distanceArray[i] == math.MaxInt32 {
			distanceArray[i] = -1
		}
	}
	return distanceArray

}
func main() {
	N := 7
	edges := [][]int{
		{0, 1, 2},
		{1, 3, 1},
		{4, 0, 3},
		{4, 2, 1},
		{5, 4, 1},
		{6, 4, 2},
		{6, 5, 3},
	}
	fmt.Println(shortestPath(N, edges, 0))
}
