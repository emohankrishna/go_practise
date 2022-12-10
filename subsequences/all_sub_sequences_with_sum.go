package main

import "fmt"

func getAllSubSequencesWithSum(arr []int, index, target int, l []int, ds [][]int) [][]int {
	if target == 0 {
		temp_l := make([]int, len(l))
		copy(temp_l, l)
		ds = append(ds, temp_l)
		return ds
	}
	if target < 0 || index == len(arr) {
		return ds
	}
	l = append(l, arr[index])
	ds = getAllSubSequencesWithSum(arr, index+1, target-arr[index], l, ds)
	l = l[:len(l)-1]
	ds = getAllSubSequencesWithSum(arr, index+1, target, l, ds)
	return ds
}

func getAllSubSequencesCountWithSum(arr []int, index, target int, l []int, count int) int {
	if target == 0 {
		return 1 + count
	}
	if target < 0 || index == len(arr) {
		return count
	}
	l = append(l, arr[index])
	count = getAllSubSequencesCountWithSum(arr, index+1, target-arr[index], l, count)
	l = l[:len(l)-1]
	count = getAllSubSequencesCountWithSum(arr, index+1, target, l, count)
	return count
}
func main() {
	arr := []int{1, 2, 2, 3, 4, 6, 7, 8}
	fmt.Println(getAllSubSequencesWithSum(arr, 0, 9, []int{}, [][]int{}))
	fmt.Println(getAllSubSequencesCountWithSum(arr, 0, 9, []int{}, 0))
}
