package main

import "fmt"

// DisjointSet represents the disjoint-set data structure.
type DisjointSet struct {
	parent []int
	rank   []int
}

// NewDisjointSet creates a new disjoint-set data structure with the specified size.
func NewDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, size),
		rank:   make([]int, size),
	}

	// Initialize each element as a separate set.
	for i := 0; i < size; i++ {
		ds.parent[i] = i
		ds.rank[i] = 0
	}

	return ds
}

// Find finds the representative/root element of the set containing the given element.
func (ds *DisjointSet) Find(element int) int {
	// Path compression optimization.
	if ds.parent[element] != element {
		ds.parent[element] = ds.Find(ds.parent[element])
	}

	return ds.parent[element]
}

// Union merges the sets containing the given two elements.
func (ds *DisjointSet) Union(element1, element2 int) {
	up_u := ds.Find(element1)
	up_v := ds.Find(element2)

	// Weighted-union heuristic.
	// up_u : ultimate parent of u
	// up_v : ultimate parent of v
	if ds.rank[up_u] < ds.rank[up_v] {
		ds.parent[up_u] = up_v
	} else if ds.rank[up_u] > ds.rank[up_v] {
		ds.parent[up_v] = up_u
	} else {
		ds.parent[up_v] = up_u
		ds.rank[up_u]++
	}
}

// Example usage
func main() {
	size := 10
	ds := NewDisjointSet(size)

	fmt.Println("Initial disjoint-set:")
	printSet(ds)

	ds.Union(1, 2)
	ds.Union(2, 3)
	ds.Union(4, 5)
	ds.Union(6, 7)
	ds.Union(5, 6)

	fmt.Println("\nDisjoint-set after unions:")
	printSet(ds)

	fmt.Println("\nRepresentatives of the elements:")
	for i := 0; i < size; i++ {
		fmt.Printf("Element %d: %d\n", i, ds.Find(i))
	}
	fmt.Println("\n after find of all elements:")
	printSet(ds)
}

// Helper function to print the disjoint-set elements and their representatives.
func printSet(ds *DisjointSet) {
	size := len(ds.parent)
	for i := 0; i < size; i++ {
		fmt.Printf("Element %d -> Parent: %d, Rank: %d\n", i, ds.parent[i], ds.rank[i])
	}
}
