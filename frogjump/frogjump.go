package main

import (
	"fmt"
	"math"
)

// minimum energy required by the frog to jump from 1st step to nth step
func minimumEnergyWithMemo(index int, H []int, Memo []int) int {
	if index == 0 {
		Memo[0] = 0
		return Memo[0]
	}
	left := minimumEnergyWithMemo(index-1, H, Memo) + int(math.Abs(float64(H[index]-H[index-1])))
	right := math.MaxInt
	if index > 1 {
		right = minimumEnergyWithMemo(index-2, H, Memo) + int(math.Abs(float64(H[index]-H[index-2])))
	}

	Memo[index] = int(math.Min(float64(left), float64(right)))
	return Memo[index]
}

func minimumEnergyWithTabular(H []int) int {
	DP := make([]int, len(H))
	DP[0] = 0
	for i := 1; i < len(H); i++ {
		left := DP[i-1] + int(math.Abs(float64(H[i]-H[i-1])))
		right := math.MaxInt
		if i > 1 {
			right = DP[i-2] + int(math.Abs(float64(H[i]-H[i-2])))
		}
		DP[i] = int(math.Min(float64(left), float64(right)))
	}
	return DP[len(H)-1]
}
func main() {

	H := []int{10, 20, 30, 10}
	N := len(H)
	Memo := make([]int, N+1)
	for key := range Memo {
		Memo[key] = -1
	}
	fmt.Println(minimumEnergyWithMemo(3, H, Memo))
}
