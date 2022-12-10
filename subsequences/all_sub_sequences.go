package main

import "fmt"

func subSequences(a []int, i int, l []int, ds [][]int) [][]int {
	if i == len(a) {
		ds = append(ds, l)
		return ds
	}
	l = append(l, a[i])
	ds = subSequences(a, i+1, l, ds)
	l = l[:len(l)-1]
	ds = subSequences(a, i+1, l, ds)
	return ds
}

func main() {
	ds := [][]int{}
	fmt.Println(subSequences([]int{1, 2, 3}, 0, []int{}, ds))
	// [[1 2 3] [1 2] [1 3] [1] [2 3] [2] [3] []]
}
