package main

import (
	"container/heap"
	"fmt"
)

type IntHeap struct {
	items []int
}

func (h IntHeap) Len() int {
	return len(h.items)
}

func (h IntHeap) Less(i, j int) bool {
	return h.items[i] > h.items[j]
}
func (h IntHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
func (h *IntHeap) Push(x any) {
	(*h).items = append((*h).items, x.(int))
}

func (h *IntHeap) Pop() any {
	old := (*h).items
	n := len(old)
	// x := (*h)[0]
	x := old[n-1]
	(*h).items = old[0 : n-1]
	return x
}
func (h *IntHeap) Size() int {
	return len((*h).items)
}
func (h *IntHeap) Peek() (int, bool) {
	if h.Size() == 0 {
		return 0, false
	}
	return (*h).items[0], true
}

func main() {
	h := &IntHeap{items: []int{21, 16, 13, 12, 9, 7, 3}}
	heap.Init(h)
	for h.Len() > 0 {
		val, _ := h.Peek()
		fmt.Printf("PeekMax: %d \n", val)
		fmt.Println("ExtractMax : ", heap.Pop(h))
	}
}
