package main

import "fmt"

type Heap struct {
	items    []int
	capacity int
}

func (heap *Heap) Init(arr []int) {
	// Convert array into Heap
}

// LeftChild index
func (heap *Heap) LeftChild(i int) int {
	return 2*i + 1
}

func (heap *Heap) RightChild(i int) int {
	return 2*i + 2
}

func (heap *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (heap *Heap) Swap(i, j int) {
	heap.items[i], heap.items[j] = heap.items[j], heap.items[i]
}

func (heap *Heap) Len() int {
	return len(heap.items)
}

func (heap *Heap) Insert(x int) {
	if heap.Len() == 0 {
		heap.items = append(heap.items, x)
		return
	}
	heap.items = append(heap.items, x)
	heap.HeapifyUp()
}

func (heap *Heap) ExtractMin() (int, bool) {
	if heap.Len() <= 0 {
		return 0, false
	}
	x := heap.items[0]
	if heap.Len() == 1 {
		heap.items = []int{}
		return x, true
	}
	heap.items[0] = heap.items[heap.Len()-1]
	heap.items = heap.items[:heap.Len()-1]
	heap.HeapifyDown(0)
	return x, true
}

func (heap *Heap) HeapifyUp() {
	index := heap.Len() - 1
	for heap.items[heap.Parent(index)] > heap.items[index] {
		heap.Swap(heap.Parent(index), index)
		index = heap.Parent(index)
	}
}
func (heap *Heap) HeapifyDown(index int) {
	if index < 0 {
		return
	}
	minIndex := index
	lIndex := heap.LeftChild(index)
	rIndex := heap.RightChild(index)
	if lIndex < heap.Len() && rIndex < heap.Len() {
		if heap.items[lIndex] < heap.items[minIndex] {
			minIndex = lIndex
		}
		if heap.items[rIndex] < heap.items[minIndex] {
			minIndex = rIndex
		}
		if index != minIndex {
			heap.Swap(index, minIndex)
			heap.HeapifyDown(minIndex)
		}
	}
}

func (heap *Heap) Print() {
	fmt.Println(heap.items)
}

func main() {
	heap := &Heap{
		items:    []int{},
		capacity: 10,
	}
	heap.Insert(21)
	heap.Insert(16)
	heap.Insert(13)
	heap.Insert(12)
	heap.Insert(9)
	heap.Insert(7)
	heap.Insert(3)
	heap.Print()
	for {
		val, ok := heap.ExtractMin()
		fmt.Println("extracted min from heap : ", val)
		if !ok {
			break
		}
	}

}
