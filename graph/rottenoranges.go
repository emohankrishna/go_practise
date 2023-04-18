package main

import "fmt"

type Tuple struct {
	x, y, time int
}

func isInBound(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	q := []Tuple{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				visited[i][j] = 2
				q = append(q, Tuple{i, j, 0})
			} else {
				// this else condition is not required as we define int slice which has default as 0
				visited[i][j] = 0
			}
		}
	}
	maxTime := 0
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, +1, 0, -1}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		maxTime = max(maxTime, t.time)
		for i := 0; i < len(drow); i++ {
			nRow := t.x + drow[i]
			nCol := t.y + dcol[i]
			if isInBound(grid, nRow, nCol) && grid[nRow][nCol] == 1 && visited[nRow][nCol] == 0 {
				q = append(q, Tuple{nRow, nCol, t.time + 1})
				visited[nRow][nCol] = 2
			}
		}

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] != 2 && grid[i][j] == 1 {
				return -1
			}
		}
	}
	return maxTime
}

func main() {
	// M, N := 3, 3
	// grid := make([][]int, M)
	// for i := 0; i < M; i++ {
	// 	grid[i] = make([]int, N)
	// }
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}
