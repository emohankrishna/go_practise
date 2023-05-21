package main

import (
	"fmt"
	"math"
)

func main() {

	var arr []int = []int{1, 2, 3, 4, 5}
	var N int = len(arr)
	x := math.Ceil(math.Log(float64(N)) / math.Log(float64(2)))
	max_size := 2*int(math.Pow(float64(2), x)) - 1
	ST := make([]int, max_size)
	construct(arr, 0, N-1, 0, ST)
	fmt.Println(ST)
}
func getMid(s, e int) int {
	return s + (e-s)/2
}
func construct(arr []int, start, end, current int, ST []int) int {
	if start == end {
		ST[current] = arr[start]
		return ST[current]
	}
	mid := getMid(start, end)
	ST[current] = construct(arr, start, mid, start*2+1, ST) + construct(arr, mid+1, end, start*2+2, ST)
	return ST[current]
}
