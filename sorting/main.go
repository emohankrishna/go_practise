package main

import "fmt"

// find the minimum index in each iteration and swap
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

// start from 0 index and 1 index
func bubbleSortI(arr []int) []int {
	n := len(arr)
	isDidSwap := true
	a := make([]int, n)
	copy(a, arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				isDidSwap = false
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		if isDidSwap {
			break
		}
		fmt.Println("run")

	}
	return a
}

func insertionSort(arr []int) []int {
	n := len(arr)
	a := make([]int, n)
	copy(a, arr)
	for i := 1; i < n-1; i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	return a
}

func insertionSortII(arr []int) []int {
	n := len(arr)
	a := make([]int, n)
	copy(a, arr)
	for i := 1; i < n-1; i++ {
		j := i
		for j > 0 && a[j-1] > a[j] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}
	return a
}

// MergeSort function recursively divides the array into smaller subarrays
// until each subarray has only one element. Then it merges and sorts the subarrays.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

// Helper function to merge two sorted arrays into a single sorted array
func merge(left, right []int) []int {
	size := len(left) + len(right)
	merged := make([]int, size)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		merged[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		merged[k] = right[j]
		j++
		k++
	}

	return merged
}

// Quicksort is a recursive function that sorts an array using the QuickSort algorithm.
// It chooses the first element as the pivot and partitions the array into two halves.
// The elements less than the pivot are placed on the left side, and the elements greater
// than the pivot are placed on the right side. The pivot is then placed in its final
// sorted position, and the same process is repeated for the two sub-arrays.
func Quicksort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIndex := partition(arr, low, high)

		// Recursively sort the left and right sub-arrays
		Quicksort(arr, low, pivotIndex-1)
		Quicksort(arr, pivotIndex+1, high)
	}
}

// Partition function rearranges the array such that all elements less than the pivot
// are placed on the left side, and all elements greater than the pivot are placed on
// the right side. It returns the final position of the pivot element.
func partition(arr []int, low, high int) int {
	pivot := arr[low] // Choose the first element as the pivot
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j {
			// Swap arr[i] and arr[j] to move elements to the correct side of the pivot
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap the pivot element with the last element in the smaller elements section
	arr[low], arr[j] = arr[j], arr[low]

	return j // Return the pivot index
}
func main() {
	// arr := []int{13, 46, 24, 52, 20, 9}
	arr := []int{9, 13, 20, 24, 46, 52}
	// selection sort
	select_sort_arr := insertionSort(arr)
	fmt.Println(select_sort_arr)
	select_sort_arr = insertionSortII(arr)
	fmt.Println(select_sort_arr)

}
