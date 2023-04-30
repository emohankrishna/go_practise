package main

import (
	"container/list"
	"fmt"
)

type Tuple struct {
	x, y int
}

func isLand(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && j < len(grid[0]) && i < len(grid) && grid[i][j] == 1
}
func numEnclaves(grid [][]int) int {
	m, n := len(grid), 0
	if m > 0 {
		n = len(grid[0])
	}
	visited := make([][]bool, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]bool, n)
	}
	q := list.New()
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 && grid[i][j] == 1 {
				q.PushBack(Tuple{i, j})
			}
		}
	}
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		tuple := e.Value.(Tuple)
		if !visited[tuple.x][tuple.y] {
			visited[tuple.x][tuple.y] = true
			//Top
			if isLand(grid, tuple.x-1, tuple.y) {
				q.PushBack(Tuple{tuple.x - 1, tuple.y})
			}
			//Bottom
			if isLand(grid, tuple.x+1, tuple.y) {
				q.PushBack(Tuple{tuple.x + 1, tuple.y})
			}
			//Right
			if isLand(grid, tuple.x, tuple.y+1) {
				q.PushBack(Tuple{tuple.x, tuple.y + 1})
			}
			//Left
			if isLand(grid, tuple.x, tuple.y-1) {
				q.PushBack(Tuple{tuple.x, tuple.y - 1})
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				count++
			}
		}
	}
	return count
}

func main() {
	// grid := [][]byte{
	// 	{1, 1, 1, 1, 0},
	// 	{1, 1, 0, 1, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 0},
	// }
	// fmt.Println(grid)
	// fmt.Println(numIslands(grid))

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(numEnclaves(image))
}
