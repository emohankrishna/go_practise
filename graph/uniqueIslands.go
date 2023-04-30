package main

import (
	"fmt"
	"strconv"
)

type Tuple struct {
	x, y int
}

func isLand(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && grid[i][j] == 1
}
func distinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	l := make([]Tuple, 0)
	M := make(map[string]struct{})
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				l, visited = dfs(grid, i, j, visited)
				fmt.Println(l, visited)
				fmt.Println()
				firstCell := l[0]
				str := ""
				for _, cell := range l {
					str = str + strconv.Itoa(cell.x-firstCell.x) + strconv.Itoa(cell.y-firstCell.y)
				}
				M[str] = struct{}{}
			}
		}
	}
	fmt.Println(M)
	return len(M)
}

func dfs(grid [][]int, i, j int, visited [][]bool) ([]Tuple, [][]bool) {
	dRow := []int{-1, 0, 1, 0}
	dCol := []int{0, 1, 0, -1}
	q := []Tuple{
		{i, j},
	}
	l := make([]Tuple, 0)
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		if !(visited[pop.x][pop.y]) {
			visited[pop.x][pop.y] = true
			l = append(l, Tuple{pop.x, pop.y})
			for k := 0; k < len(dRow); k++ {
				if isLand(grid, pop.x+dRow[k], pop.y+dCol[k]) {
					q = append(q, Tuple{pop.x + dRow[k], pop.y + dCol[k]})
				}
			}
		}
	}
	return l, visited
}
func main() {
	grid := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}
	fmt.Println(distinctIslands(grid))
}
