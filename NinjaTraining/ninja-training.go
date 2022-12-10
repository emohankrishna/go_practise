package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func ninjaTraining(day, last int, DP, points [][]int) int {
	if day == 0 {
		maxi := 0
		for i := 0; i < 3; i++ {
			if i != last {
				maxi = max(maxi, points[day][i])
			}
		}
		return maxi
	}
	if DP[day][last] != -1 {
		return DP[day][last]
	}
	maxi := 0
	for i := 0; i < 3; i++ {
		if i != last {
			maxi = max(maxi, points[day][i]+ninjaTraining(day-1, i, DP, points))
		}
	}
	DP[day][last] = maxi
	fmt.Println(DP)
	return DP[day][last]
}
func main() {
	points := [][]int{
		{10, 40, 70},
		{20, 50, 80},
		{30, 60, 90},
	}
	DP := make([][]int, 3)
	for i := 0; i < 3; i++ {
		DP[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			DP[i][j] = -1
		}
	}
	N := len(points)
	fmt.Println(ninjaTraining(N-1, 3, DP, points))
}
