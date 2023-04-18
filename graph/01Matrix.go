package main

import "fmt"

// Distance of the nearest 0 for each cell in a matrix
type Tuple struct {
	x, y, cost int
}

func isBound(g [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(g) && j < len(g[0])
}

func updateMatrix(g [][]int) [][]int {
	m, n := len(g), len(g[0])
	visited := make([][]bool, m)
	output := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		output[i] = make([]int, n)
	}
	q := make([]Tuple, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 0 {
				visited[i][j] = true
				q = append(q, Tuple{
					i, j, 0,
				})
			}
		}
	}
	dRow := []int{-1, 0, 1, 0}
	dCol := []int{0, 1, 0, -1}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		output[front.x][front.y] = front.cost
		for i := 0; i < len(dRow); i++ {
			if isBound(g, front.x+dRow[i], front.y+dCol[i]) && !visited[front.x+dRow[i]][front.y+dCol[i]] {
				visited[front.x+dRow[i]][front.y+dCol[i]] = true
				q = append(q, Tuple{front.x + dRow[i], front.y + dCol[i], front.cost + 1})
			}
		}
	}
	return output

}
func main() {
	g := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(updateMatrix(g))
}
