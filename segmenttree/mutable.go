package main

type NumArray struct {
	arr []int
	seg []int
}

func buildSegmentTree(arr []int, seg []int, low, high, curr int) int {
	if low == high {
		seg[curr] = arr[low]
		return arr[low]
	}
	mid := low + (high-low)/2
	seg[curr] = buildSegmentTree(arr, seg, low, mid, 2*curr+1) + buildSegmentTree(arr, seg, mid+1, high, 2*curr+2)
	return seg[curr]
}

func Constructor(nums []int) NumArray {
	seg := make([]int, 4*len(nums))
	buildSegmentTree(nums, seg, 0, len(nums)-1, 0)
	return NumArray{
		arr: nums,
		seg: seg,
	}
}

/*
*

	seg : Segment Tree
	low: 0
	high: len(arr)-1
	curr : 0
	index : index to which the value is updated
	val : value
*/
func UpdateUtil(seg []int, low, high, curr, index, val int) {
	if high == low {
		seg[curr] = val
		return
	}
	mid := low + (high-low)/2
	if index <= mid {
		UpdateUtil(seg, low, mid, 2*curr+1, index, val)
	} else {
		UpdateUtil(seg, mid+1, high, 2*curr+2, index, val)
	}
	seg[curr] = seg[2*curr+1] + seg[2*curr+2]
}

func (this *NumArray) Update(index int, val int) {
	this.arr[index] = val
	UpdateUtil(this.seg, 0, len(this.arr)-1, 0, index, val)
}

func SumRangeUtil(seg []int, low, high, left, right, curr int) int {

	if high < left || right < low {
		return 0
	}

	if low >= left && high <= right {
		return seg[curr]
	}
	// partial overlap
	mid := low + (high-low)/2
	return SumRangeUtil(seg, low, mid, left, right, 2*curr+1) + SumRangeUtil(seg, mid+1, high, left, right, 2*curr+2)
}

func (this *NumArray) SumRange(left int, right int) int {
	// fmt.Println(this.seg)
	return SumRangeUtil(this.seg, 0, len(this.arr)-1, left, right, 0)
}
func main() {
	sgTree := Constructor([]int{1, 3, 5})
	sgTree.Update(1, 2)
	sgTree.SumRange(0, 2) // 8
}
