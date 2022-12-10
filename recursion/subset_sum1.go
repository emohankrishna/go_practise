package main

import "fmt"

/*
*
The below function will return all the sums that produce from the combination of given array.
It also prints the combinations
*/
func subSetSum1(arr []int, sum, index int, ds []int, res []int) []int {
	if index >= len(arr) {
		// append the sum to store
		ds = append(ds, sum)
		// Print the combination which is responsible to get the sum
		fmt.Println(res)
		return ds
	}
	res = append(res, arr[index])
	ds = subSetSum1(arr, sum+arr[index], index+1, ds, res)
	res = res[:len(res)-1]
	ds = subSetSum1(arr, sum, index+1, ds, res)
	return ds
}

func subSet1_Ptr(arr []int, index int, ds []int, res *[][]int) {
	if index >= len(arr) {
		// append the sum to store
		*res = append(*res, append([]int{}, ds...))
		return
	}
	subSet1_Ptr(arr, index+1, ds, res)
	ds = append(ds, arr[index])
	subSet1_Ptr(arr, index+1, ds, res)
	return
}

func main() {
	res := [][]int{}
	subSet1_Ptr([]int{1, 2, 3, 4, 5}, 0, []int{}, &res)
	fmt.Println(res, "length : ", len(res))
}
