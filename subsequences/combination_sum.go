package main

import (
	"fmt"
	"sort"
)

func combinationSumUtil(arr, ds []int, target, index, N int, result [][]int) [][]int {
	// fmt.Printf("\nds : %v index : %v target : %v ",ds,index,target)
	// fmt.Printf("\nresult : %v  ",result)

	if index == N {
		return result
	}
	if target == 0 {
		// make sure to copy all the elements from ds to another tempds so that reference will be different.
		tempDs := make([]int, len(ds))
		copy(tempDs, ds)
		result = append(result, tempDs)
		return result
	} else if arr[index] > target {
		return result
	}
	// Pick the number at index
	result = combinationSumUtil(arr, append(ds, arr[index]), target-arr[index], index, N, result)
	// Don't Pick the number at index and move forword.
	return combinationSumUtil(arr, ds, target, index+1, N, result)
}

func combinationSum(candidates []int, target int) [][]int {
	// sort.Slice(candidates, func(i, j int) bool {
	// 	return candidates[i] < candidates[j]
	// })
	sort.Ints(candidates)
	// fmt.Println(candidates)
	return combinationSumUtil(candidates, []int{}, target, 0, len(candidates), [][]int{})
}

func main() {
	arr := []int{2, 3, 5}
	result := combinationSum(arr, 15)
	fmt.Println(result)
}
