package main

import (
	"container/list"
	"fmt"
)

/*
*
get the maximum element in each sliding window

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7      3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7

*
*/
func MaxSlidingWindow(nums []int, k int) []int {
	// it will hold the index's of the maximum element in the current sliding window
	deQue := list.New()
	result := []int{}
	for i := 0; i < k; i++ {
		//we have to remove all the indexes in the deQue if current element is greater than the elements in the deQue
		for deQue.Len() != 0 && nums[i] > nums[deQue.Back().Value.(int)] {
			deQue.Remove(deQue.Back())
		}
		deQue.PushBack(i)
	}
	for i := k; i < len(nums); i++ {
		result = append(result, nums[deQue.Front().Value.(int)])
		// remove the indexes from the deQueue which are not in the sliding window
		for deQue.Len() != 0 && deQue.Front().Value.(int) <= i-k {
			deQue.Remove(deQue.Front())
		}
		// if current element is greater than last index element remove the index from deQueue
		for deQue.Len() != 0 && nums[i] > nums[deQue.Back().Value.(int)] {
			deQue.Remove(deQue.Back())
		}
		deQue.PushBack(i)
	}
	result = append(result, nums[deQue.Front().Value.(int)])
	return result
}
func main() {
	var nums []int = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(MaxSlidingWindow(nums, k))
	fmt.Println(MaxSlidingWindow([]int{4, 3, 11}, 3))
}
