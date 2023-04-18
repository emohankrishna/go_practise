package main

import (
	"container/list"
	"fmt"
)

type Tuple struct {
	x, y int
}

func isLand(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && j < len(grid[0]) && i < len(grid) && grid[i][j] == 1
}
func numIslands(grid [][]byte) int {
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
			if grid[i][j] == 1 && !visited[i][j] {
				q.PushBack(Tuple{i, j})
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
				count++
			}

		}
	}
	return count
}
func isColorChanged(image [][]int, i, j, color int) bool {
	return image[i][j] == color
}
func isBound(image [][]int, i, j, oldColor int) bool {
	return i >= 0 && j >= 0 && i < len(image) && j < len(image[0]) && image[i][j] == oldColor
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	q := list.New()
	q.PushBack(Tuple{sr, sc})
	oldColor := image[sr][sc]
	for q.Len() > 0 {
		e := q.Front()
		t := e.Value.(Tuple)
		q.Remove(e)
		if isBound(image, t.x, t.y, oldColor) && !isColorChanged(image, t.x, t.y, color) {
			image[t.x][t.y] = color
			//Bottom
			if isBound(image, t.x+1, t.y, oldColor) && !isColorChanged(image, t.x+1, t.y, color) {
				q.PushBack(Tuple{t.x + 1, t.y})
			}
			//Top
			if isBound(image, t.x-1, t.y, oldColor) && !isColorChanged(image, t.x-1, t.y, color) {
				q.PushBack(Tuple{t.x - 1, t.y})
			}
			//Right
			if isBound(image, t.x, t.y+1, oldColor) && !isColorChanged(image, t.x, t.y+1, color) {
				q.PushBack(Tuple{t.x, t.y + 1})
			}
			//Left
			if isBound(image, t.x, t.y-1, oldColor) && !isColorChanged(image, t.x, t.y-1, color) {
				q.PushBack(Tuple{t.x, t.y - 1})
			}
		}
	}
	return image
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
	fmt.Println(floodFill(image, 1, 1, 2))
}
