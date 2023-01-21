package main

import (
	"fmt"
)

// https://leetcode.com/problems/n-queens/description/
// https://leetcode.com/problems/n-queens-ii/description/

func isValidMove(row, col, N int, grid [][]byte) bool {
	if row < 0 || row >= N || col < 0 || col >= N {
		return false
	}
	// Check the Diagonal from DOWN_RIGHT to TOP_LEFT &
	for p, q := row, col; p >= 0 && q >= 0; p, q = p-1, q-1 {
		if grid[p][q] == byte('Q') {
			return false
		}
	}
	// Check the Columns from RIGHT to LEFT
	for q := col; q >= 0; q = q - 1 {
		if grid[row][q] == byte('Q') {
			return false
		}
	}
	// Check Diagonal from TOP_LEFT to DOWN_RIGHT
	for p, q := row, col; p < N && q >= 0; p, q = p+1, q-1 {
		if grid[p][q] == byte('Q') {
			return false
		}
	}
	return true
}

// Backtracking Approach
func nQueens(col, N int, grid [][]byte, res *[][]string) {
	if col == N {
		temp := []string{}
		for j := 0; j < N; j++ {
			str := string(grid[j])
			temp = append(temp, str)
		}
		*res = append(*res, temp)
	}
	//for loop is responsible for rows
	for i := 0; i < N; i++ {
		if isValidMove(i, col, N, grid) {
			grid[i][col] = 'Q'
			// Recursion is responsible for col
			nQueens(col+1, N, grid, res)
			grid[i][col] = '.'
		}
	}
}

func main() {
	n := 4
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = byte('.')
		}
	}
	res := [][]string{}
	nQueens(0, n, grid, &res)
	fmt.Println(res)
	// Input: n = 4
	// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
}
