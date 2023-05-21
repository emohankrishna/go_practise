package main

import "fmt"

func build(seg []int, arr []int, i, low, high int) int {
	if low == high {
		seg[i] = arr[low]
		return arr[low]
	}
	mid := low + (high-low)/2
	seg[i] = build(seg, arr, 2*i+1, low, mid) + build(seg, arr, 2*i+2, mid+1, high)
	return seg[i]
}

/*
*

	low : left of segment tree
	high : right of segment tree
	l : left of range query
	r : right of range
	curr : curr position in segment tree
*/
func getSum(seg []int, low, high, l, r, cur int) int {
	// No Overlap case
	if high < l || r < low {
		return 0
	}
	// Full Overlap
	if low >= l && high <= r {
		return seg[cur]
	}
	// Partial overlap
	mid := low + (high-low)/2
	return getSum(seg, low, mid, l, r, 2*cur+1) + getSum(seg, mid+1, high, l, r, 2*cur+2)
}

func update(seg []int, low, high, cur, index, val int) {
	if high == low {
		seg[cur] = val
		return
	}
	mid := low + (high-low)/2
	if index <= mid {
		update(seg, low, mid, 2*cur+1, index, val)
	} else {
		update(seg, mid+1, high, 2*cur+2, index, val)
	}
	seg[cur] = seg[2*cur+1] + seg[2*cur+2]
}

func main() {
	arr := []int{1, 3, 5}
	seg := make([]int, 4*len(arr))
	build(seg, arr, 0, 0, len(arr)-1)
	fmt.Println(seg)
	l, r := 2, 4
	fmt.Printf("Sum of index between %d - %d : %d\n", l, r, getSum(seg, 0, len(arr)-1, l, r, 0))
	update(seg, 0, len(arr)-1, 0, 1, 2)
	fmt.Println("update ", seg)
}
