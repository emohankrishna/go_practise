package main

import "fmt"

func dfs(v int, g map[int][]int, visited []bool, stack []int) []int {
	list := g[v]
	visited[v] = true
	for _, v := range list {
		if !visited[v] {
			stack = dfs(v, g, visited, stack)
		}
	}
	stack = append(stack, v)
	return stack
}
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func topoSortUsingDFS(V int, adj map[int][]int) []int {
	visited := make([]bool, V)
	stack := []int{}
	for i := 0; i < V; i++ {
		if !visited[i] {
			stack = dfs(i, adj, visited, stack)
		}
	}
	reverse(stack)
	return stack
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
	fmt.Println(topoSortUsingDFS(len(g), g)) // [5 4 2 3 1 0]
}
