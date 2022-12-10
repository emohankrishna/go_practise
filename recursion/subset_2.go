package main

import (
	"fmt"
	"sort"
)

func subset_2(arr []int, index int, ds []int, res *[][]int) {
	temp := make([]int, len(ds))
	copy(temp, ds)
	*res = append(*res, temp)
	for i := index; i < len(arr); i++ {
		if i > index && arr[i] == arr[i-1] {
			continue
		}
		ds = append(ds, arr[i])
		subset_2(arr, i+1, ds, res)
		ds = ds[:len(ds)-1]
		// subset_11(arr, i+1, ds, res)
	}
}
func main() {
	res := [][]int{}
	arr := []int{4, 4, 4, 1, 4}
	//sort to maintain the order
	sort.Ints(arr)
	subset_2(arr, 0, []int{}, &res)
	fmt.Println(res, "length : ", len(res))

	// Input: nums = [1,2,2]
	// Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
}
