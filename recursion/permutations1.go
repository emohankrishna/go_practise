package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
In this approach we are using a flag to track which element is considered
*/
func permutationRecursion(N int, arr, l1 []int, store map[int]bool, ds [][]int) [][]int {
	if len(l1) == N {
		ds = append(ds, append([]int{}, l1...))
		return ds
	}
	for i := 0; i < N; i++ {
		_, ok := store[i]
		if !ok {
			l1 = append(l1, arr[i])
			store[i] = true
			ds = permutationRecursion(N, arr, l1, store, ds)
			l1 = l1[:len(l1)-1]
			delete(store, i)
		}
	}
	return ds
}

func permutationIteration(first int, arr []int, N int, res *[][]int) {
	if first >= N {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*res = append(*res, temp)
		return
	}
	for i := first; i < N; i++ {
		arr[first], arr[i] = arr[i], arr[first]
		permutationIteration(first+1, arr, N, res)
		arr[first], arr[i] = arr[i], arr[first]
	}
}

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
func getPermutationUtil(arr []int, ds []string, k int) []string {
	if len(arr) == 1 {
		ds = append(ds, strconv.Itoa(arr[0]))
		return ds
	}
	fact := factorial(len(arr) - 1)
	d := int(k / fact)
	r := k % fact
	ds = append(ds, strconv.Itoa(arr[d]))
	arr = append(arr[:d], arr[d+1:]...)
	return getPermutationUtil(arr, ds, r)

}
func getNthPermutation(n int, k int) string {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	m := getPermutationUtil(arr, []string{}, k-1)
	return strings.Join(m, "")
}
func main() {
	res := [][]int{}
	// fmt.Println(permutationRecursion(3, []int{1, 2, 3}, []int{}, map[int]bool{}, [][]int{}))
	nums := []int{1, 2, 3}
	permutationIteration(0, nums, len(nums), &res)
	fmt.Println(res)
	fmt.Println("length : ", len(res))
	for i := 1; i <= 24; i++ {
		fmt.Println(getNthPermutation(4, i))
	}
}
