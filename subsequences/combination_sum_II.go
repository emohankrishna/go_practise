package main

import (
	"fmt"
	"sort"
)

func combinationSumUtil(arr, ds []int, target, index, N int, result *[][]int) {
	// fmt.Printf("\nds : %v index : %v target : %v ",ds,index,target)
	// fmt.Printf("\nresult : %v  ",result)
	if target == 0 {
		// make sure to copy all the elements from ds to another tempds so that reference will be different.
		tempDs := make([]int, len(ds))
		copy(tempDs, ds)
		*result = append(*result, tempDs)
	}
	for k := index; k < len(arr); k++ {
		if k > index && arr[k] == arr[k-1] {
			continue
		}
		if arr[k] > target {
			break
		}
		ds = append(ds, arr[k])
		combinationSumUtil(arr, ds, target-arr[k], k+1, N, result)
		ds = ds[:len(ds)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	// sort.Slice(candidates, func(i, j int) bool {
	// 	return candidates[i] < candidates[j]
	// })
	sort.Ints(candidates)
	ans := [][]int{}
	combinationSumUtil(candidates, []int{}, target, 0, len(candidates), &ans)
	return ans
}

func main() {
	arr := []int{1, 1, 1, 2, 2}
	result := combinationSum(arr, 4)
	fmt.Println(result)
}
